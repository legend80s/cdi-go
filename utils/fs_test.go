package utils_test

import (
	"testing"

	"github.com/legend80s/go-change-dir/utils"
)

func TestGetDirectoryLevel(t *testing.T) {
	t.Run("testname1", func(t *testing.T) {
		input := "~/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/Resources/mr.lproj"
		expected := 10

		actual := utils.GetDiretoryLevel(input)

		if actual != expected {
			t.Errorf("Level should be %d but got %d", expected, actual)
		}
	})

	t.Run("testname2", func(t *testing.T) {
		input := "/Users/liuchuanzong/workspace/legend80s/cdi"
		expected := 5

		actual := utils.GetDiretoryLevel(input)

		if actual != expected {
			t.Errorf("Level should be %d but got %d", expected, actual)
		}
	})

	t.Run("testname3", func(t *testing.T) {
		input := "dir-lab/amuser-low-info"
		expected := 2

		actual := utils.GetDiretoryLevel(input)

		if actual != expected {
			t.Errorf("Level should be %d but got %d", expected, actual)
		}
	})
}

func TestMatch(t *testing.T) {
	var tests = []struct {
		paths          []string
		matchedPathLen int
		keyword        string
		want           string
	}{
		// level should be taken into account. always choose the least level until level diff > 2
		{
			paths: []string{
				"/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
				// level 6
				"/Users/name/workspace/6/paytm/MiniRecharge",
			},
			keyword: "mr",
			// we choose "MiniRecharge" though "mr.lproj"'s priority is higher than MiniRecharge abbr
			// but "mr.lproj" is too long. diff 11 - 6 = 5 > 2
			want:           "/Users/name/workspace/6/paytm/MiniRecharge",
			matchedPathLen: 2,
		},
		// choose "MiniRecharge" because diff 6 - 2 = 4 > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/MiniRecharge",
			matchedPathLen: 2,
		},
		// choose "MiniRecharge" because diff 6 - 3 = 3 > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/3/MiniRecharge",
			matchedPathLen: 2,
		},
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/MiniRecharge",
				"/Users/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/MiniRecharge",
			matchedPathLen: 3,
		},
		// choose "mr" because diff 6 - 4 = 2 not > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/4/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/name/workspace/6/paytm/mr",
			matchedPathLen: 2,
		},
		// choose "mr" because diff 6 - 5 = 1 not > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/4/5/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/name/workspace/6/paytm/mr",
			matchedPathLen: 2,
		},
		// choose "mr" because diff 6 - 6 = 0 not > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/4/5/6/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/name/workspace/6/paytm/mr",
			matchedPathLen: 2,
		},
		// choose "mr" because diff 6 - 7 = -1 not > 2
		{
			paths: []string{
				"/Users/name/workspace/6/paytm/mr",
				"/Users/3/4/5/6/7/MiniRecharge",
			},
			keyword:        "mr",
			want:           "/Users/name/workspace/6/paytm/mr",
			matchedPathLen: 2,
		},

		{
			paths: []string{
				// level 11
				"/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
				// level 8
				"/Users/name/workspace/alipay/paytm/7/8/MiniRecharge",
			},
			keyword: "mr",
			// we choose level 8 because 11 - 8 > 2
			want:           "/Users/name/workspace/alipay/paytm/7/8/MiniRecharge",
			matchedPathLen: 2,
		},
		{
			paths: []string{
				// level 11
				"/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
				// level 9
				"/Users/name/workspace/alipay/paytm/7/8/9/MiniRecharge",
			},
			keyword: "mr",
			// we choose level 11 because 11 - 9 not > 2
			want:           "/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
			matchedPathLen: 2,
		},

		{
			paths: []string{
				// level 11
				"/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
				// level 10
				"/Users/name/workspace/alipay/paytm/7/8/9/10/MiniRecharge",
			},
			keyword: "mr",
			// we choose level 11 because 11 - 9 not > 2
			want:           "/Users/name/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/11/mr.lproj",
			matchedPathLen: 2,
		},

		{
			paths: []string{
				"/Users/liuchuanzong/workspace/alipay/paytm/MiniRecharge",
				"/Users/liuchuanzong/workspace/alipay/pc/openmonitor-bulls/src/page/MonitorRule",
				"/Users/liuchuanzong/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/Frameworks/Electron Framework.framework/Versions/A/Resources/mr.lproj",
				"/Users/liuchuanzong/workspace/legend80s/gallery/dist/mac/UGallery.app/Contents/Resources/mr.lproj",
			},
			keyword:        "mr",
			want:           "/Users/liuchuanzong/workspace/alipay/paytm/MiniRecharge",
			matchedPathLen: 4,
		},

		// full match against base name
		{paths: []string{"xxx/balance", "xxx/mini-balance"}, keyword: "balance", want: "xxx/balance", matchedPathLen: 2},

		{
			paths: []string{
				"dir-lab/amuser-low-info",
				"dir-lab/long/ali",
				"dir-lab/long/ali-test",
				"dir-lab/long/alitest",
				"dir-lab/long/hello-ali-test",
				"dir-lab/long/hello-alitest",
				"dir-lab/long-long-long-long-long/ali",
			},
			keyword:        "ali",
			want:           "dir-lab/long/ali",
			matchedPathLen: 6,
		},

		{
			paths: []string{
				"x/amuser-low-info",
				"x/long-long-long-long-long-ali",
				"xxx/balipay-a-long-dir",
				"xxx/ali-abtest-long-long-dir",
				"xxx/balance",
			},
			keyword:        "ali",
			want:           "xxx/ali-abtest-long-long-dir",
			matchedPathLen: 3,
		},

		// inverse input
		{
			paths: []string{
				"xxx/ali-abtest-long-long-dir",
				"x/amuser-low-info",
			},
			keyword:        "ali",
			want:           "xxx/ali-abtest-long-long-dir",
			matchedPathLen: 2,
		},

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

		// path base abbr prefix - match
		{paths: []string{"test/balance-recharge-sdk", "test/socialfinance"}, keyword: "social", want: "test/socialfinance", matchedPathLen: 1},
		// path base abbr suffix - no match
		{paths: []string{"test/balance-recharge-sdk", "test/socialfinance"}, keyword: "finance", want: "", matchedPathLen: 0},

		// path base abbr prefix take priority over suffix
		{paths: []string{"test/balance-recharge-sdk", "test2/sdk-balance-recharge"}, keyword: "br", want: "", matchedPathLen: 0},

		// Contains
		{paths: []string{"test/hello", "test/helpers"}, keyword: "per", want: "", matchedPathLen: 0},
		{paths: []string{"test/hello", "test/helpers"}, keyword: "ell", want: "", matchedPathLen: 0},
		// Contains: keyword len must gt 2
		{paths: []string{"test/hello", "test/helpers"}, keyword: "er", want: "", matchedPathLen: 0},

		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "MiniRecharge", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},
		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "minirecharge", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},

		{paths: []string{"~/workspace/alipay/MiniRecharge/", "test/helpers"}, keyword: "mr", want: "~/workspace/alipay/MiniRecharge/", matchedPathLen: 1},
	}

	for _, tt := range tests {
		testname := tt.keyword

		t.Run(testname, func(t *testing.T) {
			matches := []utils.PrioritizedMatcher{}

			for _, path := range tt.paths {
				matched, priority := utils.Match(tt.keyword, path)

				if matched {
					matches = append(matches, utils.PrioritizedMatcher{path, priority})
				}
			}

			if len(matches) != 0 {
				utils.SortIntelligently(matches)

				if len(matches) != tt.matchedPathLen {
					t.Errorf("Should got exactly %d path but got %d", tt.matchedPathLen, len(matches))

					// t.Errorf("got %s, want %s", target, tt.want)
				} else {
					best := utils.GetBestMatch(matches)

					if best != tt.want {
						t.Errorf(`Input "%s" got "%s" but want "%s"`, tt.keyword, best, tt.want)
					}
				}
			} else {
				if tt.want == "" {
					// 正常
				} else {
					t.Errorf("No path matched but want: %s", tt.want)
				}
			}
		})
	}
}
