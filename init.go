package main

import (
	"log"
	"os"
)

func init()  {
	LogFile, _ := os.OpenFile("./log/test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	log.SetOutput(LogFile)
}
