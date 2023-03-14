package a

import (
	"fmt"
	"github.com/zhiyanfoo/module-a/seperate-lib1/pkg/b"
)

func F1() {
	fmt.Printf("Calling module-a/seperate-lib1/b.F1 from seperate-lib2 cmd\n")
	b.F1()
}
