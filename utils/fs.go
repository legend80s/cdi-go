package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var print = fmt.Println

type ByPriorityThenLen []PrioritizedMatcher

func (s ByPriorityThenLen) Len() int {
	return len(s)
}

func (s ByPriorityThenLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPriorityThenLen) Less(i, j int) bool {
	// v1
	// return len(s[i].Path) < len(s[j].Path)

	// v2 v3
	// return s[i].priority < s[j].priority

	iPriority := s[i].Priority
	jPriority := s[j].Priority

	// fmt.Println("s[i]", s[i], "\ns[j]", s[j])
	// fmt.Println("iPriority", iPriority, "jPriority", jPriority)

	if iPriority == jPriority {
		return len(s[i].Path) < len(s[j].Path)
	}

	iLevel := GetDiretoryLevel(s[i].Path)
	jLevel := GetDiretoryLevel(s[j].Path)

	// fmt.Println("iLevel", iLevel, "jLevel", jLevel)

	// we choose the shorter one when level too deep
	// "too deep" means level diff > 2
	DIFF := 1

	if iPriority < jPriority {
		// j is the weaker one, but it can gain power from is flatter nested level
		if jLevel+DIFF < iLevel {
			// true => less is in the front
			// so false => pick jLevel
			return false
		}

		// true => less is in the front
		// so true => pick iLevel
		return true
	}

	// true => less is in the front
	// i is the weaker one, so it can gain power from level too
	return iLevel+DIFF < jLevel
}

// func min(x int, y int) int {
// 	if x < y {
// 		return x
// 	}

// 	return y
// }

// func max(x int, y int) int {
// 	if x > y {
// 		return x
// 	}

// 	return y
// }

func GetDiretoryLevel(path string) int {
	separator := string(os.PathSeparator)

	return len(
		strings.Split(
			strings.TrimLeft(path, separator),
			separator,
		),
	)
}

type PrioritizedMatcher struct {
	Path     string
	Priority int
}

func FindBestMatch(base string, dirname string, verbose bool) string {
	var matches []PrioritizedMatcher

	// print("base:", base, "change dir:", dirname)

	walkDir(base, func(path string) bool {
		matched, priority := Match(dirname, path)

		if matched {
			// target = path
			matches = append(matches, PrioritizedMatcher{path, priority})
		}

		return matched
	})

	// print("matches:", matches)
	SortIntelligently(matches)

	if verbose {
		// print("matches after sort by priority:", "[\n  "+strings.Join(matches, "\n  ")+"\n]")
		print("matches after sort by priority:", "[\n  ", matches, "\n]")
	}

	target := ""
	if matches != nil {
		target = GetBestMatch(matches)
	}

	return target
}

func SortIntelligently(matches []PrioritizedMatcher) {
	// fmt.Println("before sorting", matches)
	sort.Sort(ByPriorityThenLen(matches))
	// fmt.Println("after sorting", matches)
}

func GetBestMatch(matches []PrioritizedMatcher) string {
	return matches[0].Path
}

var IGNORED_DIRS = [...]string{
	"node_modules",
	".",
}

// walks recursively the given directory.
// https://stackoverflow.com/questions/36713248/how-to-terminate-early-in-golang-walk
// You should return an error from your walkfunc. To make sure no real error has been returned, you can just use a known error, like io.EOF.
func walkDir(baseDir string, match func(path string) bool) {
	// https://golang.org/pkg/path/filepath/#Walk
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			shouldIgnore := IncludesFunc(IGNORED_DIRS[:], func(str string) bool {
				return strings.HasPrefix(filepath.Base(path), str)
			})

			if shouldIgnore {
				// fmt.Println(path, "ignored")

				return filepath.SkipDir
			}

			if match(path) {
				// fmt.Println(path, "matched")
				return filepath.SkipDir

				// return io.EOF
			}
		}

		return nil
	})

	// if err != nil && err != io.EOF {
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", baseDir, err)

		return
	}
}

func Match(target string, path string) (bool, int) {
	// fmt.Println(target, item)

	lowerCasedTarget := strings.ToLower(target)
	base := filepath.Base(path)
	lowerCasedBase := strings.ToLower(base)

	// Full match: `cdi balance` will jump to `xxx/balance`.
	// Nested dir supported: cdi test/mini-balance` => `path/to/test/mini-balance`.
	// Redundant `lowerCasedBase == lowerCasedTarget` is for performance.
	if lowerCasedBase == lowerCasedTarget ||
		strings.HasSuffix(
			strings.ToLower(path),
			string(os.PathSeparator)+lowerCasedTarget,
		) {
		return true, 0
	}

	// 前缀。一般记忆都是前缀，故后缀不考虑
	if strings.HasPrefix(lowerCasedBase, lowerCasedTarget) {
		return true, 1
	}

	abbr := Abbr(base)

	// fmt.Println(base, abbr)

	// 支持首字母缩写
	if abbr == lowerCasedTarget {
		return true, 2
	}

	// cdi rs => balance-recharge-sdk not /helpers
	// hasAbbr := abbr != base
	// if hasAbbr && (strings.HasPrefix(abbr, lowerCased) || strings.HasSuffix(abbr, lowerCased)) {
	// 	return true
	// }

	// 最后是完整单词包含关系
	r, _ := regexp.Compile(fmt.Sprintf("\\b%s\\b", lowerCasedTarget))

	return r.MatchString(lowerCasedBase), 3
}

type DBStruct struct {
	Workspace string            `json:"workspace"`
	Shortcuts map[string]string `json:"shortcuts"`
}

// Read and unmarshal DB
// What is the difference?
// func ReadDB(dbFilepath string, verbose bool) dbStruct {
func ReadDB(dbFilepath string, verbose bool) DBStruct {
	dat := DBStruct{
		Workspace: GetDefaultSearchDir(),
		Shortcuts: map[string]string{},
	}

	byt, err := ioutil.ReadFile(dbFilepath)

	if err != nil {
		fmt.Println("Read db/shortcuts.json failed:", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Println("len(byt)", len(byt), "string(byt)", string(byt), "string(byt) != ''", string(byt) != "")
	}

	// golang check if file is empty
	if len(byt) != 0 {
		if err := json.Unmarshal(byt, &dat); err != nil {
			fmt.Println("Unmarshal", dbFilepath, "failed:", err)
			os.Exit(1)
		}
	}

	return dat
}

// db 方案能让耗时从 1.544s 下降到 0.426s
func FindBestMatchFromDB(shortcuts map[string]string, dirname string, verbose bool) string {
	// println("dbFilepath", dbFilepath)

	if verbose {
		fmt.Println(shortcuts)
	}

	if target, ok := shortcuts[dirname]; ok {
		return target
	}

	if target, ok := shortcuts[Abbr(dirname)]; ok {
		return target
	}

	return ""
}

func SaveWorkspaceToDB(dbFilepath string, database DBStruct, newWorkspacePath string, verbose bool) (bool, error) {
	database.Workspace = newWorkspacePath

	if verbose {
		fmt.Printf("write workspace: \"%s\" to db.\n", newWorkspacePath)
	}

	return saveDB(dbFilepath, database, verbose)
}

func SaveShortcutsToDB(dbFilepath string, database DBStruct, shortcut string, path string, verbose bool) {
	db := database.Shortcuts

	db[shortcut] = path
	db[Abbr(shortcut)] = path

	if verbose {
		fmt.Printf("write shortcut: \"%s\" and path:\"%s\" to db.\n", shortcut, path)
	}

	saveDB(dbFilepath, database, verbose)
}

func saveDB(dbFilepath string, database DBStruct, verbose bool) (bool, error) {
	if verbose {
		fmt.Println("save new db", database)
	}

	byt, err := json.MarshalIndent(database, "", "  ")

	if err != nil {
		fmt.Println("db json stringify failed. db:", database, "error:", err)

		return false, err
	}

	if verbose {
		print(string(byt))
	}

	err = ioutil.WriteFile(dbFilepath, byt, 0644)

	if err != nil {
		fmt.Printf("WriteFile json to\"%s\" failed\n. json: %s\n", dbFilepath, string(byt))

		fmt.Println("error:", err)

		return false, err
	}

	return true, nil
}

var DBFilepath = genDBFilepath()

func genDBFilepath() string {
	homedir, _ := os.UserHomeDir()

	dbFilepath := path.Join(homedir, "cdi-db-shortcuts.json")

	if IsFileNotExists(dbFilepath) {
		file, err := os.Create(dbFilepath)

		if err != nil {
			panic(err)
		}

		defer file.Close()
	}

	return dbFilepath
}

// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go/22467409#22467409
func IsFileNotExists(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}

func GetDefaultSearchDir() string {
	return Normalize("~/workspace")
}

func Normalize(dir string) string {
	r := regexp.MustCompile("^~")

	home, _ := os.UserHomeDir()

	return r.ReplaceAllString(dir, home)
}
