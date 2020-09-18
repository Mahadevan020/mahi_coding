package main
import (
	"fmt"
	"runtime"
)

func main(){
	x:=2
	y:=2.99
	z:="Hello GO"
	a:=`asdw
asdas aswwww`
	fmt.Printf("Value : %v %v Type: %T %T \n",x,y,x,y)
	fmt.Printf(runtime.GOOS+"\n")
	fmt.Printf(runtime.GOARCH+"\n")
	fmt.Printf("%v",z)
	fmt.Printf("%v",a)

}
