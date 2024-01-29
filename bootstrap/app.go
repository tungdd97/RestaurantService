package bootstrap

import "restaurant-service/database/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = InitEnv()
	app.Mongo = InitMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
