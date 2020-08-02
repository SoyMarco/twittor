package bd

import (
	"context"
	"time"

	"github.com/SoyMarco/twittor/models"
)

/*BorroRelacion borra la relacion de la BD*/
func BorroRelacion(t models.Relacion) (bool, error) {
	//ESTAS 4 SIEMPRE
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
