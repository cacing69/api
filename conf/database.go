package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	datasource := "root:cacing.mysql@tcp(localhost:3306)/db_gin_develop?parseTime=true&charset=utf8"
	db, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(err.Error())
	}

	// orm.RegisterDataBase("default", "mysql", datasource, 30)
	// orm.RegisterModel(new(entity.User), new(entity.Tester))

	DB = db
}
