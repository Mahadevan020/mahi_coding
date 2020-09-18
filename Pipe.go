package main

import(
	"log"
	"os"
	"text/template"
	)
var tem *template.Template
func init(){
	tem = template.Must(template.New("").Funcs(fm).ParseFiles("new6.html"))
}
func sq(a int) int{
	return a*a
}
func double(a int)int {
	return a*2
}
func divide(a int)int{
	return a/2
}
var fm = template.FuncMap{
	"db": double,
	"sq":sq,
	"dv":divide,
}
func main(){
	x:=20
	err:= tem.ExecuteTemplate(os.Stdout,"new6.html",x)
	if err!=nil {
		log.Fatal(err)
	}
}
