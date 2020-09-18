package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)
var tpl *template.Template

var fun = template.FuncMap{
	"uc" : strings.ToUpper,
	"ft" : threeAlp,
}


func threeAlp (a string) string {
	a= strings.TrimSpace(a)
	a= a[:3]
	return a
}

func init(){
	tpl= template.Must(template.New("").Funcs(fun).ParseFiles("go.html"))
}
type sage1 struct {
	Name string
	Country string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}


func main(){
		s1 := sage1{
			Name : "Elon Musk",
			Country: "USA"}
		b2 := sage1{
			Name :"Gandhi",
			Country: "India"}
		c3:= sage1{
			Name : "Hitler",
			Country: "Germany"}
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	q := []sage1{s1,b2,c3}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage1
		Transport []car
	}{
		q,
		cars,
	}
	err :=  tpl.ExecuteTemplate(os.Stdout,"go.html",data)
	if err!= nil{
		log.Fatal(err)
	}

}