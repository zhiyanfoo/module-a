package main

import (
	"fmt"
)

func main() {
	fmt.Printf("module-a main\n")
}

func init() {
	fmt.Printf("module-a init\n")
}
