package clients

import (
	"net/http"

	"github.com/BigDwarf/sahtian/internal/model"
	"github.com/BigDwarf/sahtian/internal/service/clients"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *clients.Service
}

func NewController(service *clients.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Create(ctx echo.Context) error {
	var client model.Client
	if err := ctx.Bind(&client); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.service.CreateClient(ctx.Request().Context(), &client); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, client)
}

func (c *Controller) GetAll(ctx echo.Context) error {
	clients, err := c.service.GetAllClients(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, clients)
}
