package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	nf, err := os.Open("index2.html")
	if err != nil{
		log.Fatal("gomma ",err)
	}
	fmt.Println(nf)
}
