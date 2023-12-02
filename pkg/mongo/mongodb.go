package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bartalcorn/goponents/pkg/styles"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Clients map[string]*mongo.Client
}

// MoRepo a collection of connections
var MgoRepo MongoRepo
var clients = make(map[string]*mongo.Client)

// init initializes the Mongo packge
func init() {
	MgoRepo.Clients = clients

	mongos := os.Getenv("MONGOS")
	if mongos == "" {
		return
	}

	// this allows for multiple connect strings in a single environment variable
	// [name]<connect string>,
	parts := strings.Split(mongos, ",")
	for _, v := range parts {
		v = strings.TrimLeft(v, "[")
		vp := strings.Split(v, "]")
		connectToMongo(vp[0], vp[1])
	}
}

func connectToMongo(name string, url string) error {
	// Set client optionsedital
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error Connecting to Mongo", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("PING FAILURE", err)
	}

	fmt.Println(styles.StartBlockTop.Render(fmt.Sprintf("Connected to %s MongoDB", name)))
	clients[name] = client

	return err
}
