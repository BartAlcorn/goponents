package appconfig

import (
	"go.mongodb.org/mongo-driver/mongo"

	mgo "github.com/bartalcorn/goponents/pkg/mongo"
)

type AppConfig struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"`
	Footprints []string `json:"footprints"`
	Data       struct{} `json:"data"`
}

var AppConfigs *mongo.Collection

func init() {
	AppConfigs = mgo.MgoRepo.Clients["max"].Database("lassodb").Collection("app.config")
}
