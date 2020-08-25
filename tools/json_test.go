package tools

import (
	"Go-json/model"
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-25 08:17
* @Description:
**/

var book = model.Book{
	BookID:1,
	Title:"hello world",
	SubTitle:"golang",
	Img:"model.jpg",
	Author:"go",
	Publish:"publish",
}


func TestBookToJson(t *testing.T) {
	str, err := BookToJson(&book)
	if err != nil{
		t.Log(err)
	}
	fmt.Println(str)
}

func TestMarsh(t *testing.T) {
	str, err := book.MarshalJSON()
	if err != nil{
		t.Log(err)
	}
	fmt.Println(string(str))
}

func BenchmarkBookToJson(b *testing.B) {
	for i:=0; i<b.N; i++{
		_, err := BookToJson(&book)
		if err != nil{
			fmt.Println(err)
		}
	}
}

func BenchmarkMarshJson(b *testing.B) {
	for i:=0; i<b.N; i++{
		_, err := book.MarshalJSON()
		if err != nil{
			fmt.Println(err)
		}
	}
}