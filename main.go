package main

import (
	"github.com/jackc/pgx"
	"technopark_db/agregator"
	"technopark_db/application"
	"technopark_db/handlers"
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

	conn, _ := pgx.Connect(conf)

	var a = &application.App{
		Handler: &handlers.Handler{
			Agregator: &agregator.Agregator{
				Connection: conn,
			},
		},
	}
	a.CreateRouter()
	a.Router.Run()

}
