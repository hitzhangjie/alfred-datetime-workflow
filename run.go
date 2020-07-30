package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
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
		"2006-01-02 15:04:05.999 MST",
		"2006-01-02 15:04:05.999 -0700",
		time.RFC3339,
		time.RFC3339Nano,
		time.UnixDate,
		time.RubyDate,
		time.RFC1123,
		time.RFC1123Z,
	}

	regexpTimestamp = regexp.MustCompile(`[1-9]{1}[0-9]*`)

	phaseName = "phaseName"
	phase1st  = "first"
	phase2nd  = "second"
)

func run() {

	workflow.Var(phaseName, phase2nd)

	args := workflow.Args()
	log("run", args)

	if len(args) == 0 {
		return
	}

	input := strings.Join(args, " ")

	// 处理 now
	if input == "now" {
		processNow()
		workflow.SendFeedback()
		return
	}

	// 处理时间戳
	if regexpTimestamp.MatchString(input) {
		v, err := strconv.ParseInt(args[0], 10, 32)
		if err == nil {
			processTimestamp(time.Unix(v, 0))
			workflow.SendFeedback()
			return
		}
	}

	processTimeStr(input)
	workflow.SendFeedback()
	return

}

func log(phase string, args []string) {
	buf := &bytes.Buffer{}

	fp := filepath.Join("/Users/zhangjie/Github/alfred-datetime-workflow/test.log")
	dat, err := ioutil.ReadFile(fp)
	if err == nil && len(dat) != 0 {
		buf = bytes.NewBuffer(dat)
	}

	buf.WriteString(fmt.Sprintf("time: %v, phase:%s, args: %s\n", time.Now(), phase, strings.Join(args, " ")))
	ioutil.WriteFile(fp, buf.Bytes(), 0666)
}

func processNow() {

	now := time.Now()

	// prepend unix timestamp
	secs := fmt.Sprintf("%d", now.Unix())
	workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(icon).
		Arg(secs).
		Valid(true)

	// process all time layouts
	processTimestamp(now)
}

// process all time layouts
func processTimestamp(timestamp time.Time) {
	for _, layout := range layouts {
		v := timestamp.Format(layout)
		workflow.NewItem(v).
			Subtitle(layout).
			Icon(icon).
			Arg(v).
			Valid(true)
	}
}

func processTimeStr(timestr string) {

	timestamp := time.Time{}
	layoutMatch := ""

	for _, layout := range layouts {
		t, err := time.Parse(layout, timestr)
		if err == nil {
			timestamp = t
			layoutMatch = layout
			break
		}
	}

	// prepend unix timestamp
	secs := fmt.Sprintf("%d", timestamp.Unix())
	workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(icon).
		Arg(secs).
		Valid(true)

	for _, layout := range layouts {
		if layout == layoutMatch {
			continue
		}
		v := timestamp.Format(layout)
		workflow.NewItem(v).
			Subtitle(layout).
			Icon(icon).
			Arg(v).
			Valid(true)
	}
	return
}
