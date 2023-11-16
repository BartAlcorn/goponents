package appconfig

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	var bsonDocument bson.D
	var jsonDocument map[string]interface{}
	var temporaryBytes []byte

	findOptions := options.Find()
	// findOptions.SetLimit(2)
	cur, err := AppConfigs.Find(context.TODO(), bson.M{"type": "statusCodes"}, findOptions)
	if err != nil {
		log.Println("Error Finding documents", err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		err = cur.Decode(&bsonDocument)
		if err != nil {
			log.Println("Error decoding documents", err)
		}
		temporaryBytes, err = bson.MarshalExtJSON(bsonDocument, true, true)
		err = json.Unmarshal(temporaryBytes, &jsonDocument)
		if err != nil {
			log.Println("Error Unmarshaling documents", err)
		}
		results = append(results, jsonDocument)

	}

	if err := cur.Err(); err != nil {
		log.Println("ERROR: AppConfigs Cursor:", err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err

}
