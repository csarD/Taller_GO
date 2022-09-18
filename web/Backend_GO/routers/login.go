package routers

import (
	"encoding/json"
	"github.com/csarD/Backend_GO/bd"
	"github.com/csarD/Backend_GO/jwt"
	"github.com/csarD/Backend_GO/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario o contraseña invalido"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", http.StatusBadRequest)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario o contraseña invalido", http.StatusBadRequest)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error con el token"+err.Error(), http.StatusBadRequest)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationtime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationtime,
	})
}
