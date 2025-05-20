package app

import (
	"github.com/labstack/echo/v4"

	apiUsers "github.com/BigDwarf/sahtian/internal/server/api/users"
)

func (app *ServerApplication) registerRoutes(server *echo.Echo) {
	apiRoute := server.Group("/api/v1")

	usersRoute := apiRoute.Group("/users")

	apiUsers.RegisterController(usersRoute, app.UsersService())

}
