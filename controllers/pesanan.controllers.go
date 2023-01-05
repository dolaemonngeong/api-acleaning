package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"vp_alp/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func FetchAllOrder(c echo.Context) error {

	result, err := models.FetchAllOrder()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreOrder(c echo.Context) error {
	// c_id := c.FormValue("c_id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone := c.FormValue("phone")
	time := c.FormValue("time")
	date := c.FormValue("date")
	note := c.FormValue("note")
	t_id := c.FormValue("t_id")
	tid, err := strconv.Atoi(t_id)
	c_id := c.FormValue("c_id")
	cid, err := strconv.Atoi(c_id)
	// status := "Pending"
	// fmt.Printf(status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	var res models.Response

	v := validator.New()
	var errordata = make(map[string]string)

	err1 := v.Var(name, "required")
	if err1 != nil {
		errordata["name"] = "Name is required."
		fmt.Print("nama ksg")
	}

	err2 := v.Var(address, "required")
	if err2 != nil {
		errordata["address"] = "Address is required."
		fmt.Print("addres ksg")
	}

	err4 := v.Var(phone, "required")
	if err4 != nil {
		errordata["phone"] = "Phone is require."
		fmt.Print("phone ksg")
	}

	err5 := v.Var(time, "required")
	if err5 != nil {
		errordata["time"] = "Time is require."
		fmt.Print("time ksg")
	}

	err6 := v.Var(date, "required")
	if err6 != nil {
		errordata["date"] = "Date is require."
		fmt.Print("date ksg")
	}

	err7 := v.Var(tid, "required")
	if err7 != nil {
		errordata["tid"] = "t_id is require."
		fmt.Print("tid ksg")
	}

	err8 := v.Var(cid, "required")
	if err8 != nil {
		errordata["cid"] = "cid is require."
		fmt.Print("cid ksg")
	}

	if len(errordata) != 0 {
		fmt.Println("ada ksg")
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusInternalServerError, res)

	} else {
		fmt.Println("berhasil ")
		result, _ := models.StoreOrder(name, address, phone, time, date, note, tid, cid)
		fmt.Println("berhsl bikin")
		return c.JSON(http.StatusOK, result)
		//fmt.Println("d")
	}

}

// update data
func UpdateOrder(c echo.Context) error {
	o_id := c.FormValue("o_id")
	oid, err := strconv.Atoi(o_id)
	status := c.FormValue("status")

	// result, err := models.UpdateOrder(status, oid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	var res models.Response
	v := validator.New()
	var errordata = make(map[string]string)

	// err1 := v.Var(oid, "required")
	// if err1 != nil {
	// 	errordata["oid"] = "o_id is require."
	// }

	err2 := v.Var(status, "required")
	if err2 != nil {
		errordata["status"] = "status is require."
	}

	if len(errordata) != 0 {
		fmt.Println("a")
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusInternalServerError, res)

	} else {
		fmt.Println("b")
		result, _ := models.UpdateOrder(status, oid)
		return c.JSON(http.StatusOK, result)
	}

}

func GetTechnicianOrder(c echo.Context) error {
	t_id := c.Param("t_id")
	status := c.Param("status")

	fmt.Println(status)
	fmt.Println(t_id)
	result, err := models.GetTechnicianOrder(t_id, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetCustomerOrder(c echo.Context) error {
	c_id := c.Param("c_id")
	status := c.Param("status")

	fmt.Println(status)
	fmt.Println(c_id)
	result, err := models.GetCustomerOrder(c_id, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetOrderByID(c echo.Context) error {
	o_id := c.Param("o_id")
	oid, err := strconv.Atoi(o_id)
	fmt.Println(oid)
	result, err := models.GetOrderByID(oid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
