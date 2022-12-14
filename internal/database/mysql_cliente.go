package database

import (
	"database/sql"
	"github.com/allrivenjs/course-phones-review/internal/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)
	if err != nil {
		logs.Log().Errorf("cannot create db tentat: %v", err.Error())
		panic("..")
	}
	err = db.Ping()
	if err != nil {
		logs.Log().Warn("cannot connect to mysql")
	}
	return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()
}
