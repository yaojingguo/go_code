package main

import (
	"bytes"
	"log"
	"fmt"
	"os"
)

func exit() {
	log.Printf("my name is %s\n", "xiaoyu")
	log.Fatalln("fatal error")
	log.Println("this message should not be printed")
}

func createLogger() {

	var (
		buf bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Print(&buf)
}

func toFile() {
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix ", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")
}

func main() {
	log.Println("hello, world")
	log.Printf("faile to fetch %s with GoLang version %d", "www.google", 8)
	//exit()
	//createLogger()
	//toFile()
}
