package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SoyMarco/twittor/bd"
	"github.com/SoyMarco/twittor/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña debe ser al menos de 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al registrar Usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logró insertar Usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
