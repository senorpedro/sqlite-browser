package main

import (
	"flag"
	"log"
)

var (
	file = flag.String("db", "", "name of sqlite .db file")
)

func main() {
	flag.Parse()

	if *file == "" {
		log.Fatalf("Err: You need to provide the path to a sqlite db file with --db")
	}

	s := SqliteReader{fileName: *file}
	s.open()
	s.printTables()

}
