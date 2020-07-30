package main

import (
	aw "github.com/deanishe/awgo"
)

const (
	helpURL = "https://github.com/hitzhangjie/alfred-datetime-workflow"
)

func main() {
	workflow = aw.New(aw.HelpURL(helpURL))

	phase := workflow.Config.GetString(phaseName)

	if phase != phase2nd {
		workflow.Run(run)
		return
	}

	workflow.Run(rerun)
}
