// main.go

package main

import logger "go-training/Day2/assignment8/logger/logs"

func main() {
	loggerInstance := logger.New()
	loggerInstance.log.Println("Hello")
}
