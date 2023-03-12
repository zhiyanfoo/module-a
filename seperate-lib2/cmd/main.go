package main

import (
	"fmt"

	"github.com/zhiyanfoo/module-a/pkg/a"
	sla "github.com/zhiyanfoo/module-a/seperate-lib1/pkg/a"
)

func main() {
	fmt.Printf("seplib 2 main.go\n")
	sla.F1()
	a.F1()
}
