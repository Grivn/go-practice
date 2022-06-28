package pkga

import (
	"fmt"
	"github.com/Grivn/go-practice/howtoinit/pkgb"
)

func init() {
	fmt.Println("here is a init")
}

func MethodA() {
	fmt.Println("method a")
	pkgb.MethodB()
}
