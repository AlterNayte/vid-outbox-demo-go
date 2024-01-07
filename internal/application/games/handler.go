package games

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vid-outbox-demo-go/internal/domain"
)

type Handler struct {
	Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) GetGames(c echo.Context) error {
	results, err := h.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, results)
}

func (h *Handler) CreateGame(c echo.Context) error {
	var newGame domain.Game
	if err := c.Bind(&newGame); err != nil {
		return err
	}
	game, err := h.Create(&newGame)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, game)
}
