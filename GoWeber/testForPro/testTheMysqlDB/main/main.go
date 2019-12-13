package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"GoWeber/database/config"
)

func main() {
	var dataSourceName string
	dataSourceName = fmt.Sprintf("%s:%s@(%s:%s)/%s",
		config.Database.USER, config.Database.PASSWD, config.Database.HOST, config.Database.PORT, config.Database.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	fmt.Println("dataSourceName is:", dataSourceName)
	if err != nil {
		fmt.Println("[-] while connect mysql, there is a err:", err)
	} else {
		fmt.Println("[+] connect mysql successfully!")
	}
	defer db.Close()		// 关闭数据库连接

}
