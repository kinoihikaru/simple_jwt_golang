package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(lReq LoginRequest) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: lReq.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtsecrect)
	fmt.Println(err)

	return tokenString

}
