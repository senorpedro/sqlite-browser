package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Table struct {
	Name string
}

type SqliteReader struct {
	fileName string
	DB       *gorm.DB
}

func (s *SqliteReader) open() {
	db, err := gorm.Open(sqlite.Open(s.fileName), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	s.DB = db
}

func (s SqliteReader) printTables() {
	var tables []*Table

	s.DB.Raw(`
	SELECT name FROM sqlite_schema
	WHERE 
		type ='table' 
		AND name NOT LIKE 'sqlite_%';`).Scan(&tables)

	for _, table := range tables {

		println(table.Name)
	}

}
