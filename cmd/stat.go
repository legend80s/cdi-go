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

	fmt.Println("Search directory:", db.Workspace)
	// fmt.Println("db", db)

	if len(db.Shortcuts) == 0 {
		fmt.Println("DB empty:", db.Shortcuts)
		return
	}

	// panic: runtime error: index out of range [1] with length 1
	// table.Header = &simpletable.Header{
	// 	Cells: []*simpletable.Cell{
	// 		{Text: "db.Workspace"},
	// 	},
	// }

	for key, val := range db.Shortcuts {
		r := []*simpletable.Cell{
			{Text: key},
			{Text: val},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
