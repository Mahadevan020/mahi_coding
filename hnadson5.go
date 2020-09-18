package main

import "fmt"

type mahi int
type hotdog mahi
func main(){
	var x mahi
	var y hotdog
	fmt.Printf("value : %v Type : %T",x,x)
	x=42
	y=hotdog(x)
	fmt.Println("\n",y)
	fmt.Printf("value : %v Type : %T",y,y)
}
