package main

import (
	"testing"
)

func TestRegexp(t *testing.T) {
	v := "1596117933"
	ok := regexpTimestamp.MatchString(v)
	if !ok {
		t.Fatal("not match")
	}
	t.Logf("match")
}
