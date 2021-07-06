package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/legend80s/go-change-dir/cmd"
	"github.com/legend80s/go-change-dir/utils"
)

var println = fmt.Println

// var print = fmt.Print

func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] dirname_you_want_cd_into\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

// In nodejs I use __dirname . What is the equivalent of this in Golang?
func getRuntimeDirname(dir string) string {
	dir, err := filepath.Abs(filepath.Dir(dir))

	if err != nil {
		panic(err)
	}

	return dir
}

func main() {
	flag.Usage = myUsage
	verbose := flag.Bool("verbose", false, "show more information")
	walk := flag.Bool("walk", false, "should walk directory tree")

	// stat cmd
	statCmd := flag.NewFlagSet("stat", flag.ExitOnError)
	// set search dir cmd
	setSearchDirCmd := flag.NewFlagSet("set-root", flag.ExitOnError)
	// clear cache cmd
	clearCmd := flag.NewFlagSet("clear", flag.ExitOnError)

	switch os.Args[1] {
	case "stat":
		statCmd.Parse(os.Args[2:])
		cmd.Stat(utils.DBFilepath)
		os.Exit(0)
	case "set-root":
		if len(os.Args) != 3 {
			fmt.Println("`search-dir` required. Example: cdi set-root <search-dir>")
			os.Exit(1)
		}

		setSearchDirCmd.Parse(os.Args[2:])
		cmd.SetSearchDir(utils.DBFilepath, os.Args[2], *verbose)
		os.Exit(0)
	case "clear":
		clearCmd.Parse(os.Args[2:])
		cmd.ClearDB(utils.DBFilepath, *verbose)
		os.Exit(0)
	}

	flag.Parse()

	if *verbose {
		// https://bash.cyberciti.biz/guide/$@
		fmt.Println("os.Args", os.Args)
	}

	if flag.NArg() == 0 {
		println("No target dirname passed.")
		flag.Usage()
		os.Exit(1)
	}

	// golang如何判断是从源码运行还是从二进制文件运行?
	runtimeDirname := getRuntimeDirname(os.Args[0])

	if *verbose {
		fmt.Println("__dirname", runtimeDirname)
	}
	// 如何获取 positional arguments
	dirname := os.Args[flag.NFlag()+1]
	base := cmd.GetSearchDir()

	// println("verbose", *verbose)

	// TODO: 测试下 --walk
	target, db := "", utils.ReadDB(utils.DBFilepath, *verbose)

	if !*walk {
		target = utils.FindBestMatchFromDB(db.Shortcuts, dirname, *verbose)
		// println("target from db:", target)

		if target != "" {
			if *verbose {
				println("From DB")
			}

			cd(target)
			return
		}
	}

	target = utils.FindBestMatch(base, dirname, *verbose)

	if target != "" {
		if *verbose {
			println("From walk dir")
		}

		utils.SaveShortcutsToDB(utils.DBFilepath, db, dirname, target, *verbose)

		cd(target)
		return
	}

	fmt.Printf("no dirname as \"%s\" match found in %s\n", dirname, base)

	os.Exit(1)
}

func cd(dir string) {
	// cd /Users/liuchuanzong/workspace/alipay/infrastructure/eslint-config-paytm-cli/test/mini-balance
	// 应该找到最短的
	// 深度优先导致不行哦
	fmt.Print(dir)

	// if err := os.Chdir(dir); err != nil {
	// 	log.Fatal(err)
	// }
}
