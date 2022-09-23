package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func ParsingToken() {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFyaWFkaSIsImV4cCI6MTY2Mzc3MTgyN30.ADWWLdR_OPsCKlgqNM87pejHCEiVBd4hDyQKBUAFGEs"
	var claims = &Claims{}
	tkn, _ := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtsecrect, nil
	})
	if !tkn.Valid {
		fmt.Println("token tidak valid")
	} else {
		fmt.Sprintf("Welcome %s!", claims.Username)
		fmt.Println(claims.Username)
	}
}
