package askme

import (
	"github.com/bashmohandes/go-askme/web"
	"github.com/bashmohandes/go-askme/web/askme/controllers"
)

// App represents the AskMe web app
type App struct {
	*framework.App
}

// NewApp creates a new instance of Ask App
func NewApp(base *framework.App, hc *controllers.HomeController, pc *controllers.ProfileController) *App {
	app := &App{
		App: base,
	}

	return app
}