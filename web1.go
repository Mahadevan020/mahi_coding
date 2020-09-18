package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Mahi"

	str := fmt.Sprint( `
		<!Doctype html>
		<html lang= "eu" >
		<head>	
		<meta charset="UTF-8">
		<title>Hello World</title>
		</head>
		<body>
		<h1> ` + name + `</h1>
		</body>
		</html>
		`)
	nf, err := os.Create("index1.html")
	if err != nil {
		log.Fatal("error creating file",err)
	}
	defer nf.Close()

	io.Copy(nf,strings.NewReader(str))

}