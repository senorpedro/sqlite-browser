package main

import (
	"flag"
	"log"

	"senorpedro.com/sqlite-browser/db"
	"senorpedro.com/sqlite-browser/tui"
)

var (
	file = flag.String("db", "", "name of sqlite .db file")
)

func main() {
	flag.Parse()

	if *file == "" {
		log.Fatalf("Err: You need to provide the path to a sqlite db file with --db")
	}

	s := db.SqliteReader{}
	s.Open(*file)

	tui.StartUI(&s)
}
