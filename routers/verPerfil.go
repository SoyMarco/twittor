package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SoyMarco/twittor/bd"
)

/*VerPerfil permite extraer los valores del perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro:"+err.Error(), 400)
		return
	}
	w.Header().Set("contect-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
