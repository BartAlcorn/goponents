package acquire

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(id string) (*Acquire, error) {
	var result *Acquire
	filter := bson.M{"id": id}
	err := Acquires.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("FindOne", id, err)
	}

	return result, err

}
