package app

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func (app *ServerApplication) Database() *mongo.Client {
	if app.dbClient != nil {
		return app.dbClient
	}

	return &mongo.Client{}
}
