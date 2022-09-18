package routers

import (
	"errors"
	"github.com/csarD/Backend_GO/bd"
	"github.com/csarD/Backend_GO/models"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) == 2 {
		return claims, false, string(""), errors.New("Invalid token")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	return claims, true, string(""), err
}
