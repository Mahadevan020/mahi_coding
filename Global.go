package main

import (
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

func init(){
	tpl, _ = template.ParseFiles("new8.html")
}
type csk struct{
	Name string
	Age int
}
type mi struct{
	Name string
	Age int
}
func main(){
	/*msd := csk{
		Name: "Dhoni",
		Age:39,
	}
	chr := csk{
		Name:"Chahar",
		Age:23,
	}
	noname := csk{
		Name:"",
		Age:40,
	}
	rohit := mi{
		Name:"Rohit",
		Age:34,
	}
	bum := mi{
		Name:"Bumhbra",
		Age:24,
	}
	nn := mi{
		Name:"",
		Age:25,
	}*/
	//c:= []csk{msd,chr,noname}
	//m:= []mi{rohit,bum,nn}
	/*data := struct {
		ch []csk
		mu []mi
	}{
		c,
		m,
	}*/
	x:= struct {
		Num1 int
		Num2 int
	}{
		20,
		10,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"new8.html",x)
	if err != nil {
		log.Fatal(err)
	}
}
