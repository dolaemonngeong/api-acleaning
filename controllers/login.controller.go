package controllers

import (
	"net/http"
	// "time"
	"vp_alp/helpers"
	"vp_alp/models"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckLoginTechnician(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLoginTechnician(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Login successful", "t_id": res})
}

func CheckLoginCustomer(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLoginCustomer(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Login successful", "c_id": res})
}

// func CheckLoginTechnician1(c echo.Context) error{
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// res, err := models.CheckLoginTechnician(username, password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	// if !res {
// 	// 	return echo.ErrUnauthorized
// 	// }

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["username"] = username
// 	claims["level"] = "application"
// 	claims["exp"] = time.Now().Add(time.Hour *15).Unix()

// 	mytoken, err := token.SignedString([]byte("my-s3cr3t_k3Y"))
// 	if err != nil{
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error(),})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful", "token": mytoken})
// }

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

// func CheckLoginCustomer(c echo.Context) error{
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	res, err := models.CheckLoginCustomer(username, password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	if !res {
// 		return echo.ErrUnauthorized
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["username"] = username
// 	claims["level"] = "application"
// 	claims["exp"] = time.Now().Add(time.Hour *15).Unix()

// 	mytoken, err := token.SignedString([]byte("my-s3cr3t_k3Y"))
// 	if err != nil{
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error(),})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful", "token": mytoken})
// }
