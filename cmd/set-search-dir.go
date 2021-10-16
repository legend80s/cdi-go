package cmd

import (
	"fmt"
	"os"

	"github.com/legend80s/go-change-dir/utils"
)

func SetSearchDir(dbFilepath string, searchDir string, verbose bool) (bool, error) {
	if verbose {
		fmt.Println("new search dir", searchDir)
	}

	db := utils.ReadDB(dbFilepath, verbose)

	return utils.SaveWorkspaceToDB(dbFilepath, db, searchDir, verbose)
}

// How to Delete or Remove a File in Golang?
func ClearDB(dbFilepath string, verbose bool) {
	err := os.Remove(dbFilepath)

	if err != nil {
		fmt.Println("Remove cache", dbFilepath, "failed:", err)
	}
}

func GetSearchDir() string {
	data := utils.ReadDB(utils.DBFilepath, false)

	if data.Workspace != "" {
		return data.Workspace
	}

	workspace := utils.Normalize("~/workspace")

	if utils.IsFileNotExisting(workspace) {
		home, _ := os.UserHomeDir()

		return home
	} else {
		return workspace
	}
}
