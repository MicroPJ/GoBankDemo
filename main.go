package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/micropj/gobankdemo/cmd/gobankdemo"
)

func main() {
	myFigure := figure.NewFigure("GoBankDemo", "", true)
	myFigure.Print()
	fmt.Printf("\nPart of the Enterprise CSP\n\n")
	gobankdemo.Execute()
}
