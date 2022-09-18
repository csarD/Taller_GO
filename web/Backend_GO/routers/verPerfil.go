package routers

import (
	"encoding/json"
	"github.com/csarD/Backend_GO/bd"
	"net/http"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
