package acquire

import (
	"go.mongodb.org/mongo-driver/mongo"

	mgo "github.com/bartalcorn/goponents/pkg/mongo"
)

var Acquires *mongo.Collection

func init() {
	Acquires = mgo.MgoRepo.Clients["max"].Database("lassodb").Collection("assets.ui")
}
