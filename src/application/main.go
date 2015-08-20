package main

import (
	"net/http"
	"fmt"
	
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"

	"tasks"
	"config"
)

func main() {

	getPathFromParameterAndLoadConfigFile()

	goth.UseProviders(
		facebook.New("870850926323133", "54c9687312192961b6e2b5caa319db4b", "http://localhost:8081/auth/facebook/callback"),
	)

	gothic.GetState = func(req *http.Request) string {
		return req.URL.Query().Get("state")
	}

	router := gin.New()

	router.Use(cors.Middleware(cors.Options{
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin"},
		}))

	// Set Logger
	gin.DefaultWriter = config.GetLogFile()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())

	tasks.SetRoutes(router)

	public := router.Group("/api")

	public.GET("/", func(c *gin.Context) {
		tokenString, err := createJWTToken("AnonymousUser")
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
			return
		}
		c.JSON(200, gin.H{"accessToken": tokenString})
	})

	private := router.Group("/api/private")
	private.Use(jwt.Auth(config.TokenSecret))

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})


	authRoute := router.Group("/auth")
	authRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go to /auth/facebook"})
	})
	authRoute.GET("/facebook", func(c *gin.Context) {
		gothic.GetProviderName = getProviderFacebook
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})
	authRoute.GET("/facebook/callback", func(c *gin.Context) {

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			fmt.Fprintln(c.Writer, err)
			return
		}

		tokenString, err := createJWTToken(user.Email)
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}

		c.JSON(200, gin.H{
			"name": user.Name,
			"email": user.Email,
			"userId": user.UserID,
			"facebookAccessToken": user.AccessToken,
			"accessToken": tokenString,
		})
	})

	router.Run(config.Settings["ListenAddress"].(string))
}
