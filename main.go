package main

import (
	"github.com/jackc/pgx"
	"technopark_db/application"
)

func main() {

	conf := pgx.ConnConfig{
		User:      "sayonara",
		Password:  "boy",
		Host:      "localhost",
		Port:      5432,
		Database:  "techno",
		TLSConfig: nil,
	}
	confPool := pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: 16,
	}

	a := application.CreateApp(&confPool)
	a.CreateRouter()
	a.Router.Run()
}
