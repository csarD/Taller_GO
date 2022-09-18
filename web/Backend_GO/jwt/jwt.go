package jwt

import (
	"github.com/csarD/Backend_GO/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return "Bearer" + tokenStr, err
	}
	return "Bearer" + tokenStr, nil
}
