package cmd

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbConnection struct {  // a sql object of connection
	db		*sql.DB
}

