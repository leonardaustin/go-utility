package logextra

import (
	"log"
	"fmt"
)

func PrintErrorln(v ...interface{}) {

	log.Println("\x1b[31;1m", fmt.Sprintln(v...), "\x1b[0m")
}