package main

import (
	"fmt"
	"github.com/zhiyanfoo/module-a/lib/a"
)

func main() {
	fmt.Printf("example a main\n")
	a.F1()
}

func init() {
	fmt.Printf("example a init\n")
}
