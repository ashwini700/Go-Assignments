package logger

import (
	"log"
	"os"
)

type LoggerStore struct {
	log *log.Logger
}

func New() *LoggerStore {
	return &LoggerStore{
		log: log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
