package main

import (
	"fmt"

	aw "github.com/deanishe/awgo"
)

func main() {
	workflow := aw.New()
	workflow.Run(run)
}

func run() {
	fmt.Println("hello world")
}
