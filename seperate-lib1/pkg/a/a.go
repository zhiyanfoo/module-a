package a

import (
	"fmt"
	"github.com/zhiyanfoo/module-a/seperate-lib2/pkg/b"
)

func F1() {
	fmt.Printf("Calling module-a/seperate-lib2/b.F1 from seperate-lib1/pkg/a\n")
	b.F1()
}
