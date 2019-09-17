package main

import "fmt"

func main() {
	// 不能新增 不能删除
	var cheeses [2]string
	cheeses[0] = "marial"
	cheeses[1] = "hahaha"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses)
}
