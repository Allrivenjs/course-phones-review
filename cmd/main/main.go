package main

import (
	"github.com/allrivenjs/course-phones-review/internal/database"
	"github.com/allrivenjs/course-phones-review/internal/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationRootFolder      = "file://migrations"
	migrationsScriptsVersion = 1
	schemaName               = "phones_review"
)

func main() {
	_ = logs.InitLogger()
	client := database.NewSqlClient("root:@tcp(localhost:3306)/" + schemaName)
	doMigrate(client, schemaName)

	mux := Routes()
	server := NewServer(mux)
	server.Run()

}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Errorf("cannot create migration instance: %v", err.Error())
		return
	}
	current, _, _ := m.Version()
	logs.Log().Infof("current migration version: %v", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Infof("no changes in migration scripts")
	}

}
