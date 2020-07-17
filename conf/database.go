package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:cacing.mysql@tcp(localhost:3306)/db_source?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
