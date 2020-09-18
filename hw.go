package main

import "fmt"

func main() {
	fmt.Printf("\n Hello world it's Mahi")
	foo()
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("\n Mahi")
		}
	}
}
func foo() {
	fmt.Printf("This is foo")

}
