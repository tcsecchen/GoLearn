package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	defer file.Close()

	logger := log.New(file, "[Error]", log.LstdFlags)
	for i := 0; i < 10; i++ {
		logger.Println("这是一条日志")
		time.Sleep(5 * time.Second)
	}
}
