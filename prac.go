package main

import fmt "fmt"
var z=22
func main(){
	var a=1
	b:=2.2
	b=3
	c:=int(b)
	d:="asdw"
	fmt.Println("The value is",c,"fuck i figured it out myself",d,a,z)
	n,err :=fmt.Println()
	fmt.Println("\n",n,err)
	foo()
}

func foo(){
	fmt.Println(45/3,z)
}
