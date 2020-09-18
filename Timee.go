package main

import(
	"log"
	"os"
	"time"
	"text/template"
)
var temp *template.Template
var fm = template.FuncMap{
	"tt": timeTime,
}
func timeTime(t time.Time) string{
	return t.Format(time.Kitchen)
}
func init(){
	temp = template.Must(template.New("").Funcs(fm).ParseFiles("new5.html"))
}
func main(){
	err := temp.ExecuteTemplate(os.Stdout,"new5.html",time.Now())
	if err!= nil {
		log.Fatal(err)
	}
}
