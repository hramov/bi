package main

import (
	"fmt"
	"github.com/hramov/gvc-bi/ds/internal/storage"
	"github.com/hramov/gvc-bi/ds/internal/storage/postgres"
	"log"
)

func main() {

	source, err := storage.New("test", "postgres", postgres.Options{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "asyawear",
		SslMode:  "disable",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	row := source.QueryRow("select 1")

	if row.Err() != nil {
		log.Fatalln(row.Err().Error())
	}

	var res int

	row.Scan(&res)

	fmt.Println(res)
}
