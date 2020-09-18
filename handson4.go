package main

import "fmt"

type mahi int
func main(){
	var x mahi
	fmt.Printf("value : %v Type : %T",x,x)
	x=42
	fmt.Println("\n",x)
}