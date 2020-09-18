package printer

import (
	"fmt"
	"runtime"
)

func Hello(){
	fmt.Println(runtime.NumCPU())
}
