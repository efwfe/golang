package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type aMovie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,momitempty"`
	Actors []string
}

var movies = []aMovie{
	{Title:"aaa",Year:2019,Color:false,Actors:[]string{
		"zzz","wda",
	}},

	{Title:"cold winter",Year:2018,Color:false,Actors:[]string{
		"ad","oka",
	}},
}

func main(){
	data, err := json.MarshalIndent(movies,"","    ")
	if err != nil{
		log.Fatalf("Json Marshaling failed %s",err)
	}
	fmt.Printf("%s",data)
}