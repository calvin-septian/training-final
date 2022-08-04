package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type SQLServer struct {
	LocalDB *sql.DB
}

var DbConn SQLServer

func NewSQLConnection(conn string) *SQLServer {
	s := SQLServer{}
	db, err := sql.Open("sqlserver", conn)

	if err != nil {
		fmt.Println(err)
	}
	s.LocalDB = db

	return &s
}
