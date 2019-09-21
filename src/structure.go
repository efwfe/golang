package main

import "fmt"

// 创建结构体
type Movie struct {
	name string
	Rating float32
}

func main(){
	// var m Movie a.name a.Rating
	m := Movie{
		name:   "titanic",
		Rating: 1,
	}
	fmt.Println(m)
	fmt.Printf("%+v",m)
}
