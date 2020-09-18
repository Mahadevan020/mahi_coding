package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)
type ccars struct{
	Name string
	Rate int
}
var tp *template.Template
func init(){
	tp, _ = template.ParseFiles("new7.html")
}
func main(){
	x:= [3]int{100,20,30}
	a:=ccars{
		Name:"BMW",
		Rate:1000,
	}
	b:=ccars{
		Name:"audi",
		Rate: 2000,
	}
	c:=ccars{
		Name:"",
		Rate: 3000,
	}
	q:= []ccars{a,b,c}
	fmt.Println(q)
	/*data := struct{
		Cc []ccars
	}{
		q,

	}*/
	fmt.Print(x)
	err:=tp.ExecuteTemplate(os.Stdout,"new7.html",q)
	if err != nil{
		log.Fatal(err)
	}
}