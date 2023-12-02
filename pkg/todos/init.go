package todos

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	mgo "github.com/bartalcorn/goponents/pkg/mongo"
	"github.com/bartalcorn/goponents/pkg/styles"
	"github.com/labstack/gommon/color"
)

var Todos *mongo.Collection

func init() {
	Todos = mgo.MgoRepo.Clients["local"].Database("goponents").Collection("todos")
	fmt.Println(styles.StartBlockBottom.Render("Collection:", color.Yellow("ToDos")))
}
