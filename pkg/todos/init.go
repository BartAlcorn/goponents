package todos

import (
	"go.mongodb.org/mongo-driver/mongo"

	mgo "github.com/bartalcorn/goponents/pkg/mongo"
)

var Todos *mongo.Collection

func init() {
	Todos = mgo.MgoRepo.Clients["local"].Database("goponents").Collection("todos")
}
