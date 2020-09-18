package main

import (
	"log"
	"os"
	"text/template"
	"time"
)
var tem *template.Template
func fu (t time.Time ) string{
	return t.Format("01-02-2006")
}
var fun3= template.FuncMap{
	"dmy": fu,
}
func init(){
	tem= template.Must(template.New("").Funcs(fun3).ParseFiles("time.html"))
}
func main(){
	err:= tem.ExecuteTemplate(os.Stdout,"time.html",time.Now())
	if err!= nil{
		log.Fatal(err)
	}
}