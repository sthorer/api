package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sthorer/api/api"
	"github.com/sthorer/api/config"
)

func main() {
	conf, err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	defer conf.Client.Close()

	e := api.New(conf)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))
}
