package a

import (
	"fmt"
	"github.com/zhiyanfoo/module-a/pkg/a"
)

func F1() {
	fmt.Printf("Calling module-a/pkg/a.F3 from seperate-lib1/pkg/a\n")
	a.F3()
}
