package cmd

import "testing"

func Test_ConnectToDB(t *testing.T) {
	dbc := DbConnection{}
	err := dbc.ConnectToDB()
	if err != nil {
		t.Error("[-] Test connectToDb failed.", err.Error())
	}
	defer dbc.Close()
}

func Test_insertDataBySQLStatement(t *testing.T) {
	db, err := ConnectToDB()
	if err != nil {
		t.Error("[-] 连接数据库失败")
	}
	defer db.Close()
	result, err := insertDataBySQLStatement(db, "")
	if err != nil {
		t.Error("[-] 测试插入数据库数据失败:", err)
	}
	t.Log(result)
}
