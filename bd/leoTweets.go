package bd

import (
	"context"
	"log"
	"time"

	"github.com/SoyMarco/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil*/
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	//ESTAS 4 SIEMPRE
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets
	condicion := bson.M{
		"userid": ID,
	}
	//cuando son muchas opciones conviene usar el /bson/options
	opciones := options.Find()
	opciones.SetLimit(10)
	//Acomoda los tweeets por fecha de forma descendente por el -1
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	//Para que cada pagina salte 20 multiplicado por la pagina
	opciones.SetSkip((pagina - 1) * 10)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	//recorrer cada documento que se trajo
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		//append Agrega un [] a un elemento
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
