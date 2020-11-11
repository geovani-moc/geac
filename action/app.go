package action

import (
	"fmt"
	"os"

	"github.com/geovani-moc/geac/cmd"
	"github.com/geovani-moc/geac/entity"
)

//App strucut
type App struct {
	to       []entity.Email
	user     string
	password string
}

var _app *App

//NewApp create a new app case no exists
func NewApp() *App {
	if _app == nil {
		_app = &App{
			user:     os.Getenv("EMAIL_USER"),
			password: os.Getenv("EMAIL_PASSWORD"),
		}
	}
	return _app
}

//Run init the app
func (app *App) Run() error {
	var menu cmd.Menu
	var option int
	notExit := true

	menu.NewMenu()

	for notExit {
		menu.Print()
		fmt.Scanf("%d", &option)
		notExit = menu.Select(option)
	}

	return nil
}
