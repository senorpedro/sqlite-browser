package debug

import (
	"log"
	"os"
)

// TODO remove this
func Log(v ...any) {
	file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println(v...)

}
