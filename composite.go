package main

import (
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("new3.html","new2.html","new.html","new4.html"))
}

type sage struct {
	Name string
	Country string
}
func main(){
	s := sage{
		Name : "Elon Musk",
		Country: "USA"}
	b := sage{
		Name :"Gandhi",
		Country: "India"}
	c := sage {
		Name : "Hitler",
		Country: "Germany"}
	value := []string {"a","b","c"}
	q := []sage{s,b,c}
	val:= map[string] string{
		"India" : "Gandhi",
		"usa" : "Steve Jobs"}

	err :=  tpl.ExecuteTemplate(os.Stdout,"new3.html",value)
	if err!= nil{
		log.Fatal(err)
	}
	err =  tpl.ExecuteTemplate(os.Stdout,"new2.html",val)
	if err!= nil{
		log.Fatal(err)
	}
	err =  tpl.ExecuteTemplate(os.Stdout,"new.html",s)
	if err!= nil{
		log.Fatal(err)
	}
	err =  tpl.ExecuteTemplate(os.Stdout,"new4.html",q)
	if err!= nil{
		log.Fatal(err)
	}
}
