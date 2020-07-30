package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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
	if len(args) > 1 {
		v := strings.Join(args, " ")
		fmt.Printf("%s", v)
		return
	}

	// don't log to stdout
	//fmt.Println(args)

	ts := time.Now()
	for _, layout := range layouts {
		v := ts.Format(layout)
		workflow.NewItem(v).
			Subtitle(layout).
			Icon(icon).
			Arg(v).
			Valid(true)
	}
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
