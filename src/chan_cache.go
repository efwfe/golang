package main

import (
	"fmt"
	"time"
)

func receiver(c chan string)  {
	for msg := range c{
		fmt.Println(msg)
	}
}

func main(){
	messages := make(chan string,2)
	messages <- "hello"
	messages <- "world"
	close(messages) // 不关闭会报错 相当于不会触发完成事件
	time.Sleep(time.Second *2)
	receiver(messages)
}