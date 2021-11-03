package utils

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB // 全局的数据库操作

// init函数
func init() {
	var err error
	Db, err = sqlx.Open(`mysql`, `root:123456@(127.0.0.1:3306)/go-web-beego`)
	if err != nil {
		panic("连接错误")
		os.Exit(1)
	}
	if err = Db.Ping(); err != nil {
		panic("运行错误")
	}
}
