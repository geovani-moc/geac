package main

import (
	"log"

	"github.com/geovani-moc/geac/action"
)

func main() {
	app := action.NewApp()

	err := app.Run()
	if nil != err {
		log.Print("erro ao executar app, ", err)
	}
}
