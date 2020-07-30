package main

import (
	"testing"
)

func Test_matchedLayout(t *testing.T) {
	type args struct {
		layouts []string
		timestr string
	}
	tests := []struct {
		name       string
		args       args
		wantLayout string
		wantOk     bool
	}{
		//{"1", args{layouts, "2020-01-01"}, "", false},
		{"2", args{moreLayouts, "2020-01-01"}, "2006-01-02", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLayout, _, gotOk := matchedLayout(tt.args.layouts, tt.args.timestr)
			if gotLayout != tt.wantLayout {
				t.Errorf("matchedLayout() gotLayout = %v, want %v", gotLayout, tt.wantLayout)
			}
			if gotOk != tt.wantOk {
				t.Errorf("matchedLayout() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
