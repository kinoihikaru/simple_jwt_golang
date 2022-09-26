package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func ParsingToken(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("Authorization")
		var claims = &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtsecrect, nil
		})
		if err != nil {
			fmt.Println("token tidak valid")
			err := "Token anda tidak valid"
			w.Write([]byte(err))
			return
		}
		if !tkn.Valid {
			fmt.Println("token tidak valid")
			err := "Token anda tidak valid"
			w.Write([]byte(err))
			return
		} else {
			fmt.Sprintf("Welcome %s!", claims.Username)
			fmt.Println(claims.Username)
		}
		next(w, r, p)

	}
}
