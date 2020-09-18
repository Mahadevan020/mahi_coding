package main

import (
	"fmt"
	"path"
)
func main(){
	a,b:= path.Split("css/css1/css2/css.go")
	//var p float64 = 20.2
	var c int  = int(a[0])
	fmt.Println("dir : " + a + " , file : "+b)
	fmt.Println(c)

}
