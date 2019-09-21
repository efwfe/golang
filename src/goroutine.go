package main

import (
	"fmt"
	"time"
)

func slowFunc(){
	time.Sleep(time.Second*1)
	fmt.Println("wake up now")
}

// concurrency is not parallelism
func main(){
	go slowFunc()
	fmt.Println("haha")
	time.Sleep(time.Second*2)
}