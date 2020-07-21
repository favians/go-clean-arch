package bootstrap

import (
	"database/sql"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	App *Application
)

type Application struct {
	Config *viper.Viper
	MySql  *sql.DB
	Mongo  *mongo.Client
}

func init() {
	AppInit()
}

func AppInit() {
	App = &Application{}
	App.Config = InitConfig()
	App.MySql = InitMySqlDatabase()
	App.Mongo = InitMongoDatabase()
}
