package config

type dataBase struct {
	DB 		string
	HOST	string
	PORT	string
	USER	string
	PASSWD	string
}

var Database dataBase = dataBase{
	DB: "MYSQL",
	HOST: "192.168.163.128",
	PORT: "3306",
	USER: "root",
	PASSWD: "6",
}

