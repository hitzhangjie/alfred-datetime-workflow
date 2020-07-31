package main

import (
	aw "github.com/deanishe/awgo"
)

const (
	helpURL = "https://github.com/hitzhangjie/alfred-datetime-workflow"
)

func main() {
	workflow = aw.New(aw.HelpURL(helpURL))
	workflow.Run(run)
}
