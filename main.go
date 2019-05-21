package main

import (
	"github.com/Pickausernaame/technopark_db/application"
	"github.com/jackc/pgx"
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
	a.Router.Run(":5000")
}
