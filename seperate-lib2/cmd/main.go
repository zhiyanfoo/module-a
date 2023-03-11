package main

import (
	"fmt"

	a "github.com/zhiyanfoo/module-a/pkg/a"
	sla "github.com/zhiyanfoo/module-a/seperate-lib1/pkg/a"
)

func main() {
	sla.F1()
	a.F1()
}
