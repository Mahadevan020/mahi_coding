package main
import (
	"fmt"
)
const(
	a=42
	b=1.2
	c="asdw"
)

func main(){
	fmt.Printf("%v\n%v\n%v\n %T, %T, %T",a,b,c,a,b,c)
}