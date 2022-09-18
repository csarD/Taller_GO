package middlew

import (
	"github.com/csarD/Backend_GO/bd"
	"net/http"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Coneccion perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
