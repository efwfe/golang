package main

import (
	"fmt"
)

func res(p []int ) {
	// *p[0] error : type *[]int does not support indexing; ok for arrray not ok for slice
	for i,j := 0, len(p) -1 ;i<j; i, j = i+1,j-1{
		p[i],p[j] = p[j],p[i]
	}

}

func main(){
	a := []int{1,2,3,4}
	res(a[:])
	fmt.Println(a)
}
