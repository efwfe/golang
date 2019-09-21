package main

import (
	"fmt"
	"time"
)

func channelReader(c <- chan string){
	msg := <- c
	fmt.Println(msg)
}

func channelWriter(c chan <-string)  {
	c <- "hello world"
}

func runner(c chan string){
	for i:=1;i<10;i++{
		time.Sleep(time.Second*1)
		c<-"hello baby"
	}
}


func main()  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	stop := make(chan bool)
	go runner(channel2)
	go runner(channel1)
	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("time is out")
		stop <-true
	}()
	for{
		select {
		case msg1:= <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		case <-stop:
			return

		//case <-time.After(500 * time.Millisecond):
		//	fmt.Println("timeout giving up")
		}
	}

}