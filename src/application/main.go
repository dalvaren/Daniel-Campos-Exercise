package main

// import (
// 	"fmt"
// 	"tasks"
// )
//
// func main() {
// 	fmt.Println("This is a test package...")
// 	tasks.PrintTest()
// }

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	tokenSecret = "weAreTheChampions"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims["ID"] = "Christopher"
		token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(tokenSecret))
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
	})

	private := r.Group("/api/private")
	private.Use(jwt.Auth(tokenSecret))

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})

	r.Run(":8081")
}
