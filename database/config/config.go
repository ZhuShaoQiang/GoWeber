package config

type dataBase struct {
	DBName 		string
	HOST		string
	PORT		string
	USER		string
	PASSWD		string
}

var Database = dataBase{
	DBName: "forGo",
	HOST: "192.168.163.128",
	PORT: "3306",
	USER: "root",
	PASSWD: "6",
}
