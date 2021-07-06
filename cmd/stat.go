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
	data := utils.UnmarshalDB(dbFilepath, false)

	table := simpletable.New()

	for key, val := range data {
		r := []*simpletable.Cell{
			{Text: key},
			{Text: val},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
