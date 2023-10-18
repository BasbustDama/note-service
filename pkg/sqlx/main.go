package sqlx

import "github.com/jmoiron/sqlx"

func MustOpen(databaseDsn string) *sqlx.DB {
	conn, err := sqlx.Open("postgres", databaseDsn)
	if err != nil {
		panic(err.Error())
	}

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	return conn
}
