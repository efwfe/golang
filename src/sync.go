package main

import (
	"fmt"
	"time"
)

func slowGay(c chan string){
	t := time.NewTicker(1 * time.Second)
	// go 语言中的while 语句实现
	for {
		c <- "ping"
		<-t.C
	}
}

func main(){
	messages := make(chan string)
	go slowGay(messages)
	for{
		msg_data:= <-messages
		fmt.Println(msg_data)
	}

}
