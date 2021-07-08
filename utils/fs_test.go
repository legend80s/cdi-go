package utils_test

import (
	"sort"
	"testing"

	"github.com/legend80s/go-change-dir/utils"
)

func TestMatch(t *testing.T) {
	var tests = []struct {
		paths          []string
		matchedPathLen int
		keyword        string
		want           string
	}{
		// full match against base name
		{paths: []string{"xxx/balance", "xxx/mini-balance"}, keyword: "balance", want: "xxx/balance", matchedPathLen: 2},

		// full match against path suffix
		{paths: []string{"hello/test/mini-balance", "alipay/mini-balance"}, keyword: "test/mini-balance", want: "hello/test/mini-balance", matchedPathLen: 1},

		// full match and always match the least length
		{paths: []string{"he/test/mini-balance", "hello/test/mini-balance"}, keyword: "test/mini-balance", want: "he/test/mini-balance", matchedPathLen: 2},

		// abbr
		{paths: []string{"test/mini-balance", "test/mini-recharge"}, keyword: "mr", want: "test/mini-recharge", matchedPathLen: 1},
		// abbr 3 word
		{paths: []string{"test/balance-recharge-sdk"}, keyword: "brs", want: "test/balance-recharge-sdk", matchedPathLen: 1},

		// abbr and match the least length
		{paths: []string{"test/mini-balance", "alipay/mini-balance"}, keyword: "mb", want: "test/mini-balance", matchedPathLen: 2},
		// unless with more specific abbr
		{paths: []string{"~/test/mini-balance", "~/alipay/an-mini-balance"}, keyword: "amb", want: "~/alipay/an-mini-balance", matchedPathLen: 1},
		// abbr: path abbr not supported
		{paths: []string{"~/test/mini-balance", "~/alipay/mini-balance"}, keyword: "amb", want: "", matchedPathLen: 0},

		// path base abbr prefix
		{paths: []string{"test/balance-recharge-sdk", "test/bread"}, keyword: "br", want: "test/balance-recharge-sdk", matchedPathLen: 1},
		// path base abbr suffix
		{paths: []string{"test/balance-recharge-sdk", "test/helpers"}, keyword: "rs", want: "test/balance-recharge-sdk", matchedPathLen: 1},
		// path base abbr prefix take priority over suffix
		{paths: []string{"test/balance-recharge-sdk", "test2/sdk-balance-recharge"}, keyword: "br", want: "test/balance-recharge-sdk", matchedPathLen: 2},

		// Contains
		{paths: []string{"test/hello", "test/helpers"}, keyword: "per", want: "test/helpers", matchedPathLen: 1},
		{paths: []string{"test/hello", "test/helpers"}, keyword: "ell", want: "test/hello", matchedPathLen: 1},
		// Contains: keyword len must gt 2
		{paths: []string{"test/hello", "test/helpers"}, keyword: "er", want: "", matchedPathLen: 0},

		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "MiniRecharge", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},
		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "minirecharge", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},
		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "mr", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},
	}

	for _, tt := range tests {
		testname := tt.keyword

		t.Run(testname, func(t *testing.T) {
			matches := []string{}

			for _, path := range tt.paths {
				matched := utils.Match(tt.keyword, path)

				if matched {
					matches = append(matches, path)
				}
			}

			if len(matches) != 0 {
				sort.Sort(utils.ByLen(matches))

				if len(matches) != tt.matchedPathLen {
					t.Errorf("should got exactly %d path but got %d", tt.matchedPathLen, len(matches))

					// t.Errorf("got %s, want %s", target, tt.want)
				} else {
					if matches[0] != tt.want {
						t.Errorf("got %s, want %s", matches[0], tt.want)
					}
				}
			} else {
				if tt.want == "" {
					// 正常
				} else {
					t.Errorf("should want empty string when no path matched but got non empty string: %s", tt.want)
				}
			}
		})
	}
}
