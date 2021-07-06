package utils

import (
	"regexp"
	"strings"
)

// hello-world => hw
// hello_world => hw
// helloWorld => hw
// HelloWorld => hw
func Abbr(str string) string {
	upperRegexp := regexp.MustCompile("[A-Z]")

	if match, _ := regexp.MatchString("^[a-z]+$", str); match {
		return string(str[0])
	}

	return strings.ToLower(
		strings.Join(
			upperRegexp.FindAllString(
				strings.Title(strings.ReplaceAll(str, "_", "-")),
				-1,
			),
			""),
	)
}

func IncludesFunc(list []string, predicate func(str string) bool) bool {
	for _, val := range list {
		if predicate(val) {
			return true
		}
	}

	return false
}

type ByLen []string

func (s ByLen) Len() int {
	return len(s)
}

func (s ByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// func isLowercase(str string) bool {
// 	return strings.ToLower(str) == str
// }
