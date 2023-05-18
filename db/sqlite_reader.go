package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteReader struct {
	db         *gorm.DB
	tables     []*table
	tableNames []string
}

type table struct {
	Name string
}

type ColumnInfo struct {
	Cid     int
	Name    string
	Type    string
	NotNull int
	// DefaultValue interface{} ????
	PK int
}

func (s *SqliteReader) TableInfo(table string) []ColumnInfo {
	var columns []ColumnInfo
	// TODO fix possible SqlInjection
	s.db.Raw(fmt.Sprintf("PRAGMA table_info(%s)", table)).Scan(&columns)

	return columns
}

func (s *SqliteReader) TableContent(table string) []map[string]interface{} {

	var results []map[string]interface{}
	err := s.db.Raw(fmt.Sprintf("SELECT * FROM %s", table)).Scan(&results).Error
	if err != nil {
		panic(err)
	}

	return results
}

func (s *SqliteReader) Open(fileName string) {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	s.db = db
}

func (s *SqliteReader) readTables() {
	s.db.Raw(`
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
