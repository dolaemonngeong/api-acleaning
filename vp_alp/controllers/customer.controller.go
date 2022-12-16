package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"vp_alp/helpers"
	"vp_alp/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func FetchAllCustomer(c echo.Context) error {

	result, err := models.FetchAllCustomer()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetCustomerByUsername(c echo.Context) error {
	username := c.FormValue("username")

	result, err := models.GetCustomerByUsername(username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetCustomerByID(c echo.Context) error {
	c_id := c.Param("c_id")
	cid, err := strconv.Atoi(c_id)
	fmt.Println(cid)
	result, err := models.GetCustomerByID(cid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCustomer(c echo.Context) error {

	// c_id := c.FormValue("c_id")
	name := c.FormValue("name")
	username := c.FormValue("username")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError,
	// 		map[string]string{"message": err.Error()})
	// }

	// return c.JSON(http.StatusOK, result)

	var res models.Response
	

	v := validator.New()
	var errordata = make(map[string]string)
	err1 := v.Var(email, "required,email")
	if err1 != nil {
		errordata["email"] = "Email not valid."
	}

	err2 := v.Var(name, "required")
	if err2 != nil {
		errordata["name"] = "Name is required."
	}

	err3 := v.Var(username, "required")
	if err3 != nil {
		errordata["username"] = "Username is require."
	}

	err4 := v.Var(phone, "required")
	if err4 != nil {
		errordata["phone"] = "Phone is require."
	}

	err5 := v.Var(password, "required")
	if err5 != nil {
		errordata["password"] = "Password is require."
	}

	// password := GenerateHashPassword()

	if len(errordata) != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	} else {
		password,_ := helpers.HashPassword(password)
		result, _ := models.StoreCustomer(name, username, phone, email, password)
		return c.JSON(http.StatusOK, result)
	}

}

// update data
func UpdateCustomer(c echo.Context) error {
	c_id := c.FormValue("c_id")
	id, err := strconv.Atoi(c_id)
	name := c.FormValue("name")
	username := c.FormValue("username")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// result, err := models.UpdateCustomer(id, name, username, phone, email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	var res models.Response
	v := validator.New()
	var errordata = make(map[string]string)
	err1 := v.Var(email, "required,email")
	if err1 != nil {
		errordata["email"] = "Email not valid."
	}

	err2 := v.Var(name, "required")
	if err2 != nil {
		errordata["name"] = "Name is required."
	}

	err3 := v.Var(username, "required")
	if err3 != nil {
		errordata["username"] = "Username is require."
	}

	err4 := v.Var(phone, "required")
	if err4 != nil {
		errordata["phone"] = "Phone is require."
	}

	err5 := v.Var(password, "required")
	if err5 != nil {
		errordata["password"] = "Password is require."
	}

	if len(errordata) != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	} else {
		password,_ := helpers.HashPassword(password)
		result, _ := models.UpdateCustomer(id, name, username, phone, email, password)
		return c.JSON(http.StatusOK, result)
	}
}

// delete
func DeleteCustomer(c echo.Context) error {
	c_id := c.Param("c_id")
	id, err := strconv.Atoi(c_id)

	result, err := models.DeleteCustomer(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
