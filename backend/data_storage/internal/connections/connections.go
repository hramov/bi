package connections

import (
	"database/sql"
	"fmt"
	"sync"
)

type RawConnection struct {
	SourceId int
	DriverId int
	Dsn      string
}

type Pool map[int]*sql.DB

var pool Pool
var mu sync.RWMutex

func Connect(rc []RawConnection) []error {
	mu.Lock()
	pool = make(Pool)
	mu.Unlock()

	var err error
	var errs []error

	for _, v := range rc {
		switch v.DriverId {
		case 1:
			mu.Lock()
			pool[v.SourceId], err = connectPg(v.Dsn)
			mu.Unlock()
			if err != nil {
				errs = append(errs, err)
			}
			continue
		case 2:
			mu.Lock()
			pool[v.SourceId], err = connectSqlServer(v.Dsn)
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

func connectPg(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectSqlServer(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
