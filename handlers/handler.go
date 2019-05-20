package handlers

import (
	"github.com/Pickausernaame/technopark_db/agregator"
	"github.com/jackc/pgx"
)

type Handler struct {
	Agregator *agregator.Agregator
}

func CreateHandler(conn *pgx.ConnPoolConfig) *Handler {
	Pool, _ := pgx.NewConnPool(*conn)
	var h = &Handler{
		Agregator: &agregator.Agregator{},
	}
	h.Agregator.Connection = Pool
	return h
}

var e = map[string]string{
	"message": "Can't find user with id #42\n",
}
