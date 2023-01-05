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

func FetchAllTechnician(c echo.Context) error {

	result, err := models.FetchAllTechnician()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTechnicianByName(c echo.Context) error {
	name := c.Param("name")
	//fmt.Println(name)
	result, err := models.GetTechnicianByName(name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTechnicianByLocation(c echo.Context) error {
	k_id := c.Param("k_id")
	kid, err := strconv.Atoi(k_id)
	fmt.Println(k_id)
	result, err := models.GetTechnicianByLocation(kid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTechnicianByID(c echo.Context) error {
	t_id := c.Param("t_id")
	tid, err := strconv.Atoi(t_id)
	fmt.Println(tid)
	result, err := models.GetTechnicianByID(tid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	// var p map[string]interface{}
	// p = json.Unmarshal(result.Data, &p)
	return c.JSON(http.StatusOK, result)
	// return c.JSON(http.StatusOK, result.Data)
	// return result, nil
}

func StoreTechnician(c echo.Context) error {
	// c_id := c.FormValue("c_id")
	name := c.FormValue("name")
	username := c.FormValue("username")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	password := c.FormValue("password")
	kecamatan_id := c.FormValue("kecamatan_id")
	kid, err := strconv.Atoi(kecamatan_id)
	// status := "active"
	// fmt.Printf(status)

	// result, err := models.StoreTechnician(name, username, phone, email, password, kid)

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
		errordata["phone"] = "phone is require."
	}

	err5 := v.Var(kecamatan_id, "required")
	if err5 != nil {
		errordata["kecamatan_id"] = "kecamatan_id is require."
	}

	err6 := v.Var(password, "required")
	if err6 != nil {
		errordata["password"] = "password is require."
	}

	if len(errordata) != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	} else {
		password, _ := helpers.HashPassword(password)
		result, _ := models.StoreTechnician(name, username, phone, email, password, kid)
		return c.JSON(http.StatusOK, result)
	}

}

// update data
func UpdateTechnician(c echo.Context) error {
	t_id := c.FormValue("t_id")
	id, err := strconv.Atoi(t_id)
	name := c.FormValue("name")
	username := c.FormValue("username")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	password := c.FormValue("password")
	kecamatan_id := c.FormValue("kecamatan_id")
	kid, err := strconv.Atoi(kecamatan_id)

	// result, err := models.UpdateTechnician(id, name, username, phone, email, password, kid)

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
		errordata["phone"] = "phone is require."
	}

	err5 := v.Var(kecamatan_id, "required")
	if err5 != nil {
		errordata["kecamatan_id"] = "kecamatan_id is require."
	}

	if len(errordata) != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	} else {
		password, _ := helpers.HashPassword(password)
		result, _ := models.UpdateTechnician(id, name, username, phone, email, password, kid)
		return c.JSON(http.StatusOK, result)
	}
}
func UpdateTechnicianRate(c echo.Context) error {
	fmt.Print("update contr")
	t_id := c.FormValue("t_id")
	tid, err := strconv.Atoi(t_id)
	rate := c.FormValue("rate")
	r, err := strconv.Atoi(rate)
	// result, err := models.UpdateTechnicianRate(r, tid)

	if err != nil {
		fmt.Println("update contr eror")
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	var res models.Response
	v := validator.New()
	var errordata = make(map[string]string)

	err1 := v.Var(rate, "required")
	if err1 != nil {
		errordata["rate"] = "rate is require."
		fmt.Println("rate krg")
	}

	err2 := v.Var(tid, "required")
	if err2 != nil {
		errordata["tid"] = "t_id is require."
		fmt.Println("t_id krg")
	}

	if len(errordata) != 0 {
		fmt.Print("eror updaterate")
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	} else {
		fmt.Println("berhsl updaterate")
		result, _ := models.UpdateTechnicianRate(r, tid)
		return c.JSON(http.StatusOK, result)
	}
}

// delete
func DeleteTechnician(c echo.Context) error {
	t_id := c.Param("t_id")
	id, err := strconv.Atoi(t_id)

	result, err := models.DeleteTechnician(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
