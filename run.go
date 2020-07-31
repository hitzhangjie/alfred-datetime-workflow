package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
		time.RFC1123Z,
	}

	moreLayouts = []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999",
	}

	regexpTimestamp = regexp.MustCompile(`^[1-9]{1}\d+$`)
)

func run() {

	var err error

	args := workflow.Args()
	log("run, args: %v", args)

	if len(args) == 0 {
		return
	}

	defer func() {
		if err == nil {
			workflow.SendFeedback()
			return
		}
	}()

	// 处理 now
	input := strings.Join(args, " ")
	if input == "now" {
		processNow()
		return
	}

	// 处理时间戳
	if regexpTimestamp.MatchString(input) {
		v, e := strconv.ParseInt(args[0], 10, 32)
		if e == nil {
			processTimestamp(time.Unix(v, 0))
			return
		}
		err = e
		return
	}

	// 处理时间字符串
	err = processTimeStr(input)
	if err != nil {
		log("process time str error: %v", err)
	}
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

func processTimeStr(timestr string) error {

	timestamp := time.Time{}
	layoutMatch := ""

	layoutMatch, timestamp, ok := matchedLayout(layouts, timestr)
	if !ok {
		log("layouts not mached, timestr: %s", timestr)
		layoutMatch, timestamp, ok = matchedLayout(moreLayouts, timestr)
		if !ok {
			log("morelayouts not mached, timestr: %s", timestr)
			return errors.New("no matched time layout found")
		}
	}

	// prepend unix timestamp
	secs := fmt.Sprintf("%d", timestamp.Unix())
	workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(icon).
		Arg(secs).
		Valid(true)

	// other time layouts
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

	return nil
}

func matchedLayout(layouts []string, timestr string) (matched string, timestamp time.Time, ok bool) {

	log("layouts length: %d", len(layouts))

	for _, layout := range layouts {
		log("parse time: %s in layout: %s", timestr, layout)
		v, err := time.Parse(layout, timestr)
		if err == nil {
			return layout, v, true
		}
	}
	return
}

// 设置环境变量LOGGING_ENABLED=1/true，来激活调试日志
func log(format string, args ...interface{}) {

	v := os.Getenv("LOGGING_ENABLED")
	enabled, err := strconv.ParseBool(v)
	if err != nil || !enabled {
		return
	}

	buf := &bytes.Buffer{}

	exe, _ := os.Executable()
	dir, _ := filepath.Split(exe)
	fp := filepath.Join(dir, "awgo.log")

	dat, err := ioutil.ReadFile(fp)
	if err == nil && len(dat) != 0 {
		buf = bytes.NewBuffer(dat)
	}
	format = fmt.Sprintf("time: %v, %s\n", time.Now(), format)
	buf.WriteString(fmt.Sprintf(format, args))
	ioutil.WriteFile(fp, buf.Bytes(), 0666)
}
