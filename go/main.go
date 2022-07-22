package main

import (
	"database/sql"
	"log"

	"github.com/cyanrad/university/api"
	db "github.com/cyanrad/university/db/sqlc"
	"github.com/cyanrad/university/util"
)

func main() {
	// >> loading configuration from env file/variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't read config: ", err)
	}

	// >> connecting to the db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// >> init store and db
	store := db.NewStore(conn)
	server := api.NewServer(store)

	// >> starting API server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
