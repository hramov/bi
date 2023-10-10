package connections

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
)

type ConnectOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type ConnectResponse struct {
	Ok bool
	Id string
}

type CheckResponse struct {
	Ok bool
}

type QueryOptions struct {
	Sql string
}

type QueryResponse struct {
	Result string
}

type DataSource interface {
	Connect(options ConnectOptions) ConnectResponse
	Check() CheckResponse
	Query(options QueryOptions) QueryResponse
}

type RawConnection struct {
	SourceId int
	DriverId int
	Dsn      string
}

type Pool map[int]*sql.DB

var pool Pool
var mu sync.RWMutex

func Connect(ctx context.Context, rc []RawConnection) []error {
	mu.Lock()
	pool = make(Pool)
	mu.Unlock()

	var err error
	var errs []error

	for _, v := range rc {
		switch v.DriverId {
		case 1:
			mu.Lock()
			pool[v.SourceId], err = connectPg(ctx, v.Dsn)
			mu.Unlock()
			if err != nil {
				errs = append(errs, err)
			}
			continue
		case 2:
			mu.Lock()
			pool[v.SourceId], err = connectSqlServer(ctx, v.Dsn)
			mu.Unlock()
			if err != nil {
				errs = append(errs, err)
			}
			continue
		}
	}
	return errs
}

func Get(driver int) (*sql.DB, error) {
	mu.RLock()
	defer mu.RUnlock()
	d, ok := pool[driver]
	if !ok {
		return nil, fmt.Errorf("cannot find driver")
	}
	return d, nil
}

func connectPg(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres (%s): %v", dsn, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectSqlServer(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
