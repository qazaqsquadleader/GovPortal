package migration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"govportal/pkg/logger"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ConfigDb struct {
	Driver string
	Path   string
	Name   string
}

func NewConfDb(l *logger.Logger) *ConfigDb {
	var AppConfig ConfigDb
	raw, err := ioutil.ReadFile("configDB.json")
	if err != nil {
		l.Fatal(err.Error())
	}
	json.Unmarshal(raw, &AppConfig)
	return &AppConfig
}

func (c *ConfigDb) InitDatabase(l *logger.Logger) *sql.DB {
	db, err := sql.Open(c.Driver, c.Name)
	if err != nil {
		l.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		l.Fatal(err.Error())
	}
	return db
}

func (c *ConfigDb) CreateTables(db *sql.DB, l *logger.Logger) {
	file, err := ioutil.ReadFile("migration/db.sql")
	if err != nil {
		log.Println(err.Error())
		l.Fatal(err.Error())
	}
	fmt.Println(string(file))
	if _, err := db.Exec(string(file)); err != nil {
		log.Println(err.Error())
		l.Fatal(err.Error())
	}
}
