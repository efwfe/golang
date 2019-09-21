package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main()  {
	a := 123
	fmt.Println(reflect.TypeOf(a))
	// 类型转寒
	b := strconv.Itoa(a)
	fmt.Println(reflect.TypeOf(b))
	// 缓冲区拼接字符串
	var buffer bytes.Buffer
	for i:=0;i<500;i++{
		buffer.WriteByte('z')
	}
	fmt.Println(buffer.String())
	var str = " many spaces at here "
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(str[:5])
	fmt.Println(strings.Index(str,"at"))
}
