package main

import (
	"fmt"
	"strings"
)

func rerun() {
	log("rerun, args: %v", workflow.Args())
	fmt.Printf(strings.Join(workflow.Args(), " "))
}
