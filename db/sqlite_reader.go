package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Table struct {
	Name string
}

type SqliteReader struct {
	DB         *gorm.DB
	tables     []*Table
	tableNames []string
}

func (s *SqliteReader) Open(fileName string) {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	s.DB = db
}

func (s *SqliteReader) readTables() {

	s.DB.Raw(`
	SELECT name FROM sqlite_schema
	WHERE 
		type ='table' 
		AND name NOT LIKE 'sqlite_%';`).Scan(&s.tables)

	s.tableNames = make([]string, len(s.tables))
	for i, table := range s.tables {
		s.tableNames[i] = table.Name
	}

}

func (s SqliteReader) TableNames() []string {
	s.readTables()
	return s.tableNames
}

func (s *SqliteReader) PrintTables() {
	s.readTables()

	for _, table := range s.tableNames {
		println(table)
	}

}
