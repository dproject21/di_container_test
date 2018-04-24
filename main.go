package main

import (
	"github.com/dproject21/di_container_test/lib"
)

func main() {

	lib.CreateContainer()
	p := lib.GetPrinter()
	p.Print()
}
