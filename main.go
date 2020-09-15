package main

import (
	"Go-json/cmd"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-15 08:51
* @Description:
**/

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}