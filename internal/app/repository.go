package app

import (
	"github.com/BigDwarf/sahtian/internal/repository"
)

const (
	usersCollection = "users"
)

func (app *ServerApplication) UsersRepository() *repository.UsersRepository {
	if app.usersRepository != nil {
		return app.usersRepository
	}

	collection := app.Database().Database(app.conf.Database.Database).Collection(usersCollection)

	app.usersRepository = repository.NewUsersRepository(collection)

	return app.usersRepository
}
