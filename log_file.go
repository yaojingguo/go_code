package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("to-file.log", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println("first log message!")
}
