package action

import (
	"fmt"
	"log"

	"github.com/geovani-moc/geac/util"

	"github.com/geovani-moc/geac/cmd"
	"github.com/geovani-moc/geac/entity"
)

//App strucut
type App struct {
	To      []entity.Email
	Subject string
	Config  *util.Configuration
}

var _app *App

//NewApp create a new app case no exists
func NewApp() *App {
	var err error
	if _app == nil {
		_app = &App{
			Subject: getSubject(),
			//Bobdy: getBody(),
		}
		_app.Config, err = util.LoadConfig(".config")
		if nil != err {
			log.Fatal(err)
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

func getSubject() string {
	var subject string
	fmt.Scan(&subject)

	return subject
}
