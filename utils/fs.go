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

func FindBestMatch(base string, dirname string, verbose bool) string {
	matched := false
	target := ""
	var matches []string

	// print("base:", base, "change dir:", dirname)

	walkDir(base, func(path string) bool {
		matched = Match(dirname, path)

		if matched {
			// target = path
			matches = append(matches, path)
		}

		return matched
	})

	// print("matches:", matches)
	sort.Sort(ByLen(matches))

	if verbose {
		print("matches after sort by length:", "[\n  "+strings.Join(matches, "\n  ")+"\n]")
	}

	if matches != nil {
		target = matches[0]
	}

	return target
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

func Match(target string, path string) bool {
	// fmt.Println(target, item)

	lowerCased := strings.ToLower(target)
	base := filepath.Base(path)

	// 优先全匹配 cdi balance 将跳入 xxx/balance
	// 支持层级 cdi test/mini-balance
	if strings.HasSuffix(strings.ToLower(path), string(os.PathSeparator)+lowerCased) {
		return true
	}

	abbr := Abbr(base)

	// 支持首字母缩写
	if abbr == lowerCased {
		return true
	}

	// cdi rs => balance-recharge-sdk not /helpers
	hasAbbr := abbr != base

	if hasAbbr && (strings.HasPrefix(abbr, lowerCased) || strings.HasSuffix(abbr, lowerCased)) {
		return true
	}

	// 然后是包含关系
	if len(lowerCased) > 2 && strings.Contains(strings.ToLower(base), lowerCased) {
		return true
	}

	return false
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
		fmt.Println("len(byt)", len(byt), "string(byt)", string(byt), "string(byt) != ", string(byt) != "")
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

	if isFileNotExists(dbFilepath) {
		file, err := os.Create(dbFilepath)

		if err != nil {
			panic(err)
		}

		defer file.Close()
	}

	return dbFilepath
}

// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go/22467409#22467409
func isFileNotExists(path string) bool {
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
