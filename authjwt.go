package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var secretKey = []byte("YTY2NTJiYzBiMjAyY2FiOTJlMTM5MGIx")

// User struct to decode json,xml,form, or query data sent in the POST
type User struct {
	UserName string `json:"username" xml:"username" form:"username" query:"username"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

// Login is the /login endpoint that does the work
func Login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	username := u.UserName
	password := u.Password

	// Actually check in database (crypt password first)
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "JasonKumpf"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Accessible is an endpoint that does not require authentication
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Restricted is an enpoint that requires authentication
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// Here we pass along to whatever should happen
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

/*
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", Login)

	// Unauthenticated route
	e.GET("/", Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT(secretKey))
	r.GET("", Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
*/
