package main
import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main(){
	file,err := ioutil.ReadFile("foo.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%s",file)

	// 自定义错误
	err2 := errors.New("my error")
	fmt.Println(err2)


}