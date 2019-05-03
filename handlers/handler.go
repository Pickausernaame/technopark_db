package handlers

import "technopark_db/agregator"

type Handler struct {
	Agregator *agregator.Agregator
}

var e = map[string]string{
	"message": "Can't find user with id #42\n",
}
