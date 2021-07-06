package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
		matched = match(dirname, path)

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

func match(target string, path string) bool {
	// fmt.Println(target, item)

	lowerCased := strings.ToLower(target)
	base := filepath.Base(path)

	// 优先全匹配 cdi balance 将跳入 xxx/balance
	// 支持层级 cdi test/mini-balance
	if strings.HasSuffix(strings.ToLower(path), string(os.PathSeparator)+lowerCased) {
		return true
	}

	// 支持首字母缩写
	if Abbr(base) == lowerCased {
		return true
	}

	// 然后是包含关系
	if strings.Contains(strings.ToLower(base), lowerCased) {
		return true
	}

	return false
}

// db 方案能让耗时从 1.544s 下降到 0.426s
func FindBestMatchFromDB(dbFilepath string, dirname string, verbose bool) (string, map[string]string) {
	// println("dbFilepath", dbFilepath)

	var dat = make(map[string]string)

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

	if verbose {
		fmt.Println(dat)
	}

	if target, ok := dat[dirname]; ok {
		return target, dat
	}

	if target, ok := dat[Abbr(dirname)]; ok {
		return target, dat
	}

	return "", dat
}

func SaveToDB(dbFilepath string, db map[string]string, shortcut string, path string, verbose bool) {
	db[shortcut] = path
	db[Abbr(shortcut)] = path

	if verbose {
		fmt.Printf("write shortcut: \"%s\" and path:\"%s\" to db.\n", shortcut, path)
		fmt.Println("new db", db)
	}

	byt, err := json.MarshalIndent(db, "", "  ")

	if err != nil {
		fmt.Println("db json stringify failed. db:", db, "error:", err)

		return
	}

	if verbose {
		print(string(byt))
	}

	err = ioutil.WriteFile(dbFilepath, byt, 0644)

	if err != nil {
		fmt.Printf("WriteFile json to\"%s\" failed\n. json: %s\n", dbFilepath, string(byt))

		fmt.Println("error:", err)

		return
	}
}

func GenDBFilepath() string {
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
