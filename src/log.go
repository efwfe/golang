package main

import (
	"errors"
	"log"
	"os"
)

func main(){
	f,err := os.OpenFile("exam.log",os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Print("THis is a log message")
	var errFatal = errors.New("We only just started and we are crashing")
	log.Fatal(errFatal)
}
