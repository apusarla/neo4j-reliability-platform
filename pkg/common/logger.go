package common

import (
	"log"
	"os"
)

func InitLogger(serviceName string) {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[" + serviceName + "] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
