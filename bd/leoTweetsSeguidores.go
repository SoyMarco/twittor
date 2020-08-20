package bd

import (
	"context"
	"time"

	"github.com/SoyMarco/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	//ESTAS 4 SIEMPRE
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 10
	condiciones := make([]bson.M, 0)
	//$match busca el usuario id
	//$lookup Une 2 tablas
	//from es con que tabla se une
	//"as" es el nombramiento de la nueva tabla, se puede llamar igual
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	//unwind es para poder procesar los datos de las tablas unidas
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	//condiciona que ordene por fecha
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	//primero skip y luego limit para delimitar solo mostrar 10 tweets por vez
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 10})

	//Aggregate recorre automaticamente
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
