package users

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/BigDwarf/sahtian/internal/model"
)

type Service interface {
	GetExistingUser(ctx context.Context, id, requestedId int64) (*model.User, error)
}

const moduleName = "users"

type Controller struct {
	service Service
}

func New(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func RegisterController(route *echo.Group, service Service) {
	ctl := New(service)

	route.POST("/getUser", ctl.handlerGetUser)
}
