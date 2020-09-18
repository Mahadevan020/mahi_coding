package main

import "fmt"

func main(){
	name:="Mahi"

	tpl:=`
		<!Doctype html>
		<html lang= "eu" >
		<head>	
		<meta charset="UTF-8">
		<title>Hello World</title>
		</head>
		<body>
		<h1> ` +name+ `</h1>
		</body>
		</html>
		`
	fmt.Println(tpl)
}
