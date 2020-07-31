package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SoyMarco/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoTweet braba el Tweet en la BD*/
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	//ESTAS 4 SIEMPRE
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	//Devuelve el ultimo campo insertado
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
