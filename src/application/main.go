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
	"github.com/tommy351/gin-cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	// "github.com/itsjamie/gin-cors"

	"tasks"

	"net/http"
	"time"
	"fmt"
)

var (
	tokenSecret = "weAreTheChampions"
)

func main() {

	goth.UseProviders(
		facebook.New("870850926323133", "54c9687312192961b6e2b5caa319db4b", "http://localhost:8081/auth/facebook/callback"),
	)
	// Assign the GetState function variable so we can return the
	// state string we want to get back at the end of the oauth process.
	// Only works with facebook and gplus providers.
	gothic.GetState = func(req *http.Request) string {
		// Get the state string from the query parameters.
		return req.URL.Query().Get("state")
	}

	r := gin.New()

	r.Use(cors.Middleware(cors.Options{
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin"},
		}))

  // Global middlewares
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

	tasks.SetRoutes(r)

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		tokenString, err := createJWTToken("Christopher")
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
			return
		}
		c.JSON(200, gin.H{"accessToken": tokenString})
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


	authRoute := r.Group("/auth")
	authRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go to /auth/facebook"})
	})
	authRoute.GET("/facebook", func(c *gin.Context) {
		gothic.GetProviderName = getProviderFacebook
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})
	authRoute.GET("/facebook/callback", func(c *gin.Context) {
		// http://localhost:8081/auth/facebook/callback?code=AQDjOFmELl95lsJo5iL9K-LEpBNabKF_B4p-BRkYF7-ELIT0I-bfi_Re_0BwDfYSj-BVcvM_L4YJyGYXsootc6mWNO51njOfjmU0kXAKjYG3ptRqOqYR3j5eoF1rtZeMWbo74cJGyDmVHAyq_RGaP6jIj4HA7pDWoV2UMnxnBR_QrA0BX7m4tA0tqmGWSfnb1pMCQBaU9-L4McFXx1P-l3lC2rmjED8we4C3csuIKxBuL6U5D-_Vsy2yDn1yfdFku3yAlBLDfQ7KVH24WOn9n9kHJExHRgVps-8RUDzcThMNyAsVcYlIn0Rycvq98wRBbseDzgNXrAoUKdNvMJP5loVP#_=_

		// print our state string to the console
		fmt.Println(gothic.GetState(c.Request))

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			fmt.Fprintln(c.Writer, err)
			return
		}

		tokenString, err := createJWTToken("Christopher")
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

	r.Run(":8081")
}

func getProviderFacebook(req *http.Request) (string, error) {
	return "facebook", nil
}

func createJWTToken(userID string) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["ID"] = userID
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(tokenSecret))
	return tokenString, err
}
