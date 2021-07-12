package utils_test

import (
	"testing"

	"github.com/legend80s/go-change-dir/utils"
)

func TestAbbr(t *testing.T) {
	var tests = []struct {
		original string
		want     string
	}{
		{original: "hello-world", want: "hw"},
		{original: "hello_world", want: "hw"},
		{original: "helloWorld", want: "hw"},
		{original: "HelloWorld", want: "hw"},
		{original: "mr", want: "mr"},
		{original: "helpers", want: "helpers"},
	}

	for _, tt := range tests {
		testname := tt.original

		t.Run(testname, func(t *testing.T) {
			abbr := utils.Abbr(tt.original)

			if abbr != tt.want {
				t.Errorf("got %s, want %s", abbr, tt.want)
			}
		})
	}
}
