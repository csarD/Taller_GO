package routers

import (
	"encoding/json"
	"github.com/csarD/Backend_GO/bd"
	"github.com/csarD/Backend_GO/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Sin email", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password menor a 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
