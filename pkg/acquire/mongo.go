package acquire

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func SetTitle(a *Acquire) string {
	result := "--No Title--"

	for _, v := range a.Asset.Titles {
		result = v.Full
	}

	return result
}

func Find(id string) (*Acquire, error) {
	var result *Acquire
	// item := make(map[string]interface{})

	filter := bson.M{"id": id}
	err := Acquires.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("FindOne", id, err)
	}

	result.Title = SetTitle(result)

	return result, err
}
