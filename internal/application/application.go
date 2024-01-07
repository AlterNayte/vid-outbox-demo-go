package application

import (
	"github.com/labstack/echo/v4"
	"vid-outbox-demo-go/internal/application/games"
	"vid-outbox-demo-go/internal/persistence"
)

type Application struct {
	*echo.Echo
	Games games.Service
}

func New(conn persistence.Connection) *Application {
	return &Application{
		Echo:  echo.New(),
		Games: games.NewService(conn),
	}
}

func (a *Application) RegisterAPIRoutes() {
	g := a.Group("/api")
	gamesHandler := games.NewHandler(a.Games)
	g.GET("/games", gamesHandler.GetGames)
	g.POST("/games", gamesHandler.CreateGame)
}
