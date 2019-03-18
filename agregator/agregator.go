package agregator

import "github.com/jackc/pgx"

type Agregator struct {
	Connection *pgx.Conn
}
