package middlew

import (
	"net/http"

	"github.com/SoyMarco/twittor/bd"
)

/*ChequeoDB es el middlewere que permite conocer el estado de la BD*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
