package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("abrakadabra ")
	log.SetFlags(log.Lmsgprefix | log.Flags())

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	log.SetOutput(file)

	errorLogger := log.New(os.Stdout, "[error] ", log.Flags())

	infoLogger := log.New(os.Stdout, "[info] ", log.Flags())

	log.Println("log message")

	errorLogger.Println("error msg")
	infoLogger.Println("info msg")
}
