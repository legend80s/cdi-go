package cmd

import (
	"fmt"

	"github.com/alexeyco/simpletable"

	"github.com/legend80s/go-change-dir/utils"
)

var println = fmt.Println

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Stat(dbFilepath string) {
	db := utils.ReadDB(dbFilepath, false)

	table := simpletable.New()

	// fmt.Println("Search directory:", db.Workspace)
	// fmt.Println("db", db)

	if len(db.Shortcuts) == 0 {
		fmt.Println("DB empty:", db.Shortcuts)
		return
	}

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "Shortcut"},
			{Text: "Full Path", Align: simpletable.AlignCenter},
		},
	}

	for key, val := range db.Shortcuts {
		r := []*simpletable.Cell{
			{Text: key},
			{Text: val},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Text: "Search Root"},
			{Text: db.Workspace},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
