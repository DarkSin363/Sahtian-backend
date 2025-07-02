package app

import (
	"github.com/BigDwarf/sahtian/internal/log"
	"github.com/BigDwarf/sahtian/internal/repository"
	"github.com/BigDwarf/sahtian/internal/service/auth"
	"github.com/BigDwarf/sahtian/internal/service/clients"
	"github.com/BigDwarf/sahtian/internal/service/telegram"
	"github.com/BigDwarf/sahtian/internal/service/users"
)

func (app *ServerApplication) UsersService() *users.Service {
	if app.usersService != nil {
		return app.usersService
	}

	app.usersService = users.NewUsersService(app.UsersRepository(), app.Database(), app.conf.Storage.Bucket)

	return app.usersService
}

func (app *ServerApplication) AuthService() *auth.Service {
	if app.authService != nil {
		return app.authService
	}

	app.authService = auth.NewService(app.UsersRepository(), app.conf.Telegram.Token,
		app.conf.Debug.Enabled, app.conf.Debug.UserId)

	return app.authService
}

func (app *ServerApplication) TelegramService() *telegram.Service {
	if app.telegramService != nil {
		return app.telegramService
	}

	var err error

	app.telegramService, err = telegram.NewService(
		app.conf.Telegram.Token,
		app.conf.Telegram.AppUrl)
	if err != nil {
		log.Fatalf("Failed to create telegram service: %v", err)
	}

	return app.telegramService
}

func (app *ServerApplication) ClientService() *clients.Service {
	if app.clientService != nil {
		return app.clientService
	}

	db := app.Database().Database(app.conf.Database.Database)
	repo := repository.NewClientRepository(db)
	app.clientService = clients.NewService(repo)

	return app.clientService
}
