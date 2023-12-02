package todos

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TodosHandler struct {
	repo *mongo.Collection
}
