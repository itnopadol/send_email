package ctrl

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var dbc *sqlx.DB
var headerKeys = make(map[string]interface{})

func setHeader() {

	headerKeys = map[string]interface{}{
		"Server":                       "smart_pump_invoice",
		"Host":                         "nopadol.net:6000",
		"Content_Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token",
	}
}


func init(){
	dbc = ConnectSQL()
}

func ConnectSQL()(msdb *sqlx.DB){
	db_host := "192.168.0.7"
	db_name := "bcnp"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if (msdb.Ping()!=nil){
		fmt.Println("Error")
	}

	//fmt.Println("msdb = ",msdb.DriverName())
	return msdb
}
