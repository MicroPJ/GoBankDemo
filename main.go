package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/micropj/gobankdemo/cmd/gobankdemo"
)

func main() {
	myFigure := figure.NewFigure("gobankdemo", "", true)
	myFigure.Print()
	fmt.Printf("\n\n")
	gobankdemo.Execute()
}
