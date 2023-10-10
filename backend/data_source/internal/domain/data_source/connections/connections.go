package connections

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"plugin"
	"sync"
)

type DataSourcePlugin interface {
	Connect(ctx context.Context, dsn string) (*sql.DB, error)
}

type RawConnection struct {
	SourceId   int
	DriverId   int
	PluginName string
	Dsn        string
}

type Pool map[int]*sql.DB

type SourcePlugin map[int]string

var pool Pool
var mu sync.RWMutex

func Connect(ctx context.Context, rc []RawConnection) []error {
	mu.Lock()
	pool = make(Pool)
	mu.Unlock()

	var err error
	var errs []error

	for _, v := range rc {
		mu.Lock()
		pool[v.SourceId], err = connect(ctx, v.PluginName, v.Dsn)
		mu.Unlock()
		if err != nil {
			errs = append(errs, err)
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

func getDataSourceFromPlugin(name string) (DataSourcePlugin, error) {
	p, err := plugin.Open("plugins/data_source/" + name + "/" + name + ".so")
	if err != nil {
		return nil, fmt.Errorf("cannot open plugin %s: %v", name, err)
	}

	newFuncRaw, err := p.Lookup("New")
	if err != nil {
		log.Fatalln(err)
	}

	newFunc, ok := newFuncRaw.(func() DataSourcePlugin)
	if !ok {
		return nil, fmt.Errorf("cannot cast newFunc\n")
	}

	rawDs := newFunc()

	ds, ok := rawDs.(DataSourcePlugin)
	if !ok {
		return nil, fmt.Errorf("cannot cast interface\n")
	}

	return ds, nil
}

func connect(ctx context.Context, name, dsn string) (*sql.DB, error) {
	ds, err := getDataSourceFromPlugin(name)
	if err != nil {
		return nil, fmt.Errorf("cannot get data source: %v", err)
	}

	conn, err := ds.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to data source: %v", err)
	}

	return conn, nil
}
