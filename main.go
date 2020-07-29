package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	aw "github.com/deanishe/awgo"
)

var (
	workflow *aw.Workflow

	icon = &aw.Icon{
		Value: aw.IconClock.Value,
		Type:  aw.IconClock.Type,
	}

	layouts = []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		//time.RFC822,
		//time.RFC822Z,
		//time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		//time.Kitchen,
		//time.Stamp,
		//time.StampMilli,
		//time.StampMicro,
		//time.StampNano,
	}
)

const (
	helpURL = "https://github.com/hitzhangjie/alfred-datetime-workflow"
)

func main() {
	workflow = aw.New(aw.HelpURL(helpURL))
	workflow.Run(run)
}

func run() {

	// 获取参数
	args := workflow.Args()
	// don't log to stdout
	//fmt.Println(args)

	workflow.NewItem("this is a result").Valid(true).
		Arg(args[0]).
		Title("this is title").
		Subtitle("this is subtitle").Icon(icon)
	workflow.SendFeedback()

	buf := &bytes.Buffer{}

	fp := filepath.Join("/Users/zhangjie/Github/alfred-datetime-workflow/test.log")
	dat, err := ioutil.ReadFile(fp)
	if err == nil && len(dat) != 0 {
		buf = bytes.NewBuffer(dat)
	}

	buf.WriteString(fmt.Sprintf("time: %v, args: %v\n", time.Now(), args))
	ioutil.WriteFile(fp, buf.Bytes(), 0666)
}
