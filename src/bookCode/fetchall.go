package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)
func fetch(url string, ch chan <- string){
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err) //发送到chan通道中
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄漏资源
	if err != nil{
		ch <- fmt.Sprintf("while reading %s :%v",url,err)
		return
	}

	secs := time.Since(start)
	ch <- fmt.Sprintf("%.2fs z%7d %s",secs, nbytes, url)
}


func main(){
	start := time.Now()
	ch :=  make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url, ch) // goroutine
	}

	for range os.Args[1:]{
		fmt.Println(<-ch) // 从通道ch接受
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())

}
