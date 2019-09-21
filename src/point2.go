package main

import (
	"fmt"
)

type Drinka struct {
	Name string
	Ice bool
}

func main(){

	a := Drinka{
		Name:"zhang",
		Ice:true,

	}
	b := &a // 副本机制
	fmt.Printf("%+v \n",a)
	b.Ice = false
	fmt.Printf("%+v \n",*b)
	fmt.Printf("%+v \n",a)
	fmt.Printf("%p \n",&a)
	fmt.Printf("%p \n",b)
}