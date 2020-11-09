package main

import (
	"fmt"

	"github.com/geovani-moc/geac/cmd"
)

func main() {
	var menu cmd.Menu
	var option int
	menu.NewMenu()

	for {
		menu.Print()
		fmt.Scanf("%d", &option)
		menu.Select(option)
	}
}
