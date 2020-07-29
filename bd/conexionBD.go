package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la DB*/
var MongoCN = ConectarBD()

//HrZcSSZmjKDA2T2b
var clientOptions = options.Client().ApplyURI("mongodb+srv://marco:HrZcSSZmjKDA2T2b@cluster0.sl4yn.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar a la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión EXITOSA con la BD")
	return client
}

/*ChequeoConnection es el ping de la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
