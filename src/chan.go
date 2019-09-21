package main

import (
	"fmt"
	"time"
)

func slowFunca(c chan string){
	time.Sleep(time.Second*3)
	c <- "okay"
}

func main(){
	c := make(chan string)
	go slowFunca(c)
	fmt.Println("run 1")
	msg := <-c
	fmt.Println(msg)
	fmt.Println("run 2")
}
