package test

import (
	"github.com/syhv-git/minifier"
	"os"
	"testing"
)

func TestMinifier(t *testing.T) {
	err := minifier.Minifier("output.js", "text/javascript", "test_sample.js")
	if err != nil {
		t.Fatal(err.Error())
	}
	info, err := os.Stat("output.js")
	if err != nil {
		t.Fatal(err.Error())
	}
	if info.Size() <= 0 {
		t.Error("Output file has no contents")
	}
	os.Remove("output.js")
}
