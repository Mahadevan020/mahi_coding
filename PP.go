package main

import (
	"log"
	"text/template"
	"os"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("new2.html"))
}

func main(){
	err := tpl.Execute(os.Stdout,69)
	if err != nil {
		log.Fatal(err)
	}

}
