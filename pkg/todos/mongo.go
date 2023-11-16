package todos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert(item Todo) error {
	_, err := Todos.InsertOne(context.Background(), item)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("INSERTED: TODO: ", item.Task)

	return nil
}

func GetAll() ([]*Todo, error) {
	var results []*Todo

	findOptions := options.Find()
	// findOptions.SetLimit(2)
	cur, err := Todos.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println("ERROR: TODOS Find:", err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("ERROR: TODOS Cursor:", err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}

func GetById(id string) (*Todo, error) {
	var result *Todo
	filter := bson.M{"id": id}
	err := Todos.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("FindOne", id, err)
	}

	return result, err
}

func UpdateById(id string, todo *Todo) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"id":   todo.ID,
			"task": todo.Task,
			"done": todo.Done,
		},
	}

	result, err := Todos.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("UpdateOne", id, err)
	}
	return result, err
}

func DeleteById(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"id": id}

	result, err := Todos.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("UpdateOne", id, err)
	}
	return result, err
}
