package app

import (
	"log"
)

func Logger(data ...interface{}) {
	log.Println(data)
}
