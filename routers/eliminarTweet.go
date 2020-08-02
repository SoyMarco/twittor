package routers

import (
	"net/http"

	"github.com/SoyMarco/twittor/bd"
)

/*EliminarTweet permite borrar un Tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el paramatro ID", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	//se agrega linea de más por si se quiere agregar mas cosas
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
