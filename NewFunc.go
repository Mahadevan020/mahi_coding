package main

import (
	"log"
	"os"
	"text/template"
)
type cars struct{
	Name string
	Rate int
}
func fun1 (a string) string{
	a= a[:3]
	return a
}
func init(){
	tpl= template.Must(template.New("").Funcs(fun1).ParseFiles("go1.html"))
}

func main(){
	a:= cars {
		Name: "bmw",
		Rate: 1000}
	b:= cars{
		Name:"Audi",
		Rate:2000}
	c:= cars{
		Name: "Maruthi",
		Rate: 3000}
	cc := []cars{a,b,c}
	data:= struct {
		car []cars
	}{
		cc}
	err :=  tpl.ExecuteTemplate(os.Stdout,"go1.html",data)
	if err!= nil{
		log.Fatal(err)
	}

}

