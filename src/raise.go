package main

import (
	"fmt"

)

func raise(num int)(int,error){
	if num ==0{
		return -1,fmt.Errorf("erorrs aaa")
	}
	return 1,nil
}

func main(){
	n,err :=raise(0)
	fmt.Println(n,err)
	// error happen not stop progress
	n1,err1 :=raise(1)
	fmt.Println(n1,err1)

	//panic will stop immediately
	panic("stop now")
	n2,err2 :=raise(1)
	fmt.Println(n2,err2)

}
