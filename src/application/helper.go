package main

import (
  "os"
  "config"
	"net/http"
  "time"
	jwt_lib "github.com/dgrijalva/jwt-go"
)

func getPathFromParameterAndLoadConfigFile() {
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[0]
	}
	config.LoadConfig(path)
}

func getProviderFacebook(req *http.Request) (string, error) {
	return "facebook", nil
}

// createJWTToken generates the JWT token to be added to Request Headers
func createJWTToken(userID string) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["ID"] = userID
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	return tokenString, err
}
