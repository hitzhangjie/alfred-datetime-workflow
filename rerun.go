package main

import (
	"fmt"
	"strings"
)

func rerun() {
	log("rerun", workflow.Args())
	fmt.Printf(strings.Join(workflow.Args(), " "))
}
