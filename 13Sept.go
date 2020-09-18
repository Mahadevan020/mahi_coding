package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)
var tp *template.Template

var fun2= template.FuncMap{
	"uc1": strings.ToUpper,
	"ft1":pp,
	"ft2":kk,
}
func pp(a string) string{
	a=a[:3]
	return a
}
func kk (a int) int{
	a=a*3
	return a
}

type cars1 struct{
	Name string
	Rate int
}
type bike struct{
	Name string
	Rate int
}

func init(){
	tp= template.Must(template.New("").Funcs(fun2).ParseFiles("html5.html"))
}

func main() {
	a := cars1{
		Name: "BMW",
		Rate: 1000,
	}
	b:= cars1{
		Name:"Audi",
		Rate: 2000,
	}
	c:= bike{
		Name: "RE Bullet",
		Rate:1000,
	}
	d:=bike{
		Name:"Rx 100",
		Rate: 2000,
	}
	q:= []cars1{a,b}
	s:= []bike{c,d}

	data := struct {
		Zuk []cars1
		Xuk []bike
	}{
		q,
		s}
	err:= tp.ExecuteTemplate(os.Stdout,"html5.html",data)
	if err != nil {
		log.Fatal(err)
	}
}