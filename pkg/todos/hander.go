package todos

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TodosHandler struct {
	repo *mongo.Collection
}

// TODO unify handler and utilize the idiomatic GO Handler pattern
