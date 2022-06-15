package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDb() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=challengemeli.database.windows.net;user id=challengeuser;password=MeLi2022;database=challengeML")
	return
}
