package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"GoWeber/database/config"
)

// connect to mysql database, return two values db and err.
// if success to connect mysql, return db object and nil,
// if fail, return nil and err
func (dbc *DbConnection)ConnectToDB() (error) {
	// format the string for connect mySQLdb
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.USER, config.Database.PASSWD, config.Database.HOST, config.Database.PORT, config.Database.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	} else {
		err = db.Ping()  // test the connection between you and sql
		if err != nil {
			return err
		}
		dbc.db = db
		return nil
	}
}

func (dbc *DbConnection)insertDataBySQLStatement(db *sql.DB,sql string) (sql.Result, error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err		// if failed, return -1 and error
	}
	result, err := stmt.Exec(
		"INSERT INTO people(id, name, age)values(?, ?, ?)",
		"1",
		"zhushaoqiang",
		"21",
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dbc *DbConnection)InsertData()  {
	
}

func (dbc *DbConnection) Close() error {
	err := dbc.db.Close()  // close this connection
	if err != nil {
		return err
	}
	return nil
}



