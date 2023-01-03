package models

import (
	"fmt"
	"net/http"
	"strconv"

	// "vp_alp/controllers"
	"vp_alp/db"
)

type Customer struct {
	C_id     int    `json:"c_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email`
	Password string `json:"password" validate:"required,passwd`
	Status   string `json:"status" validate:"required`
}

// read all
func FetchAllCustomer() (Response, error) {
	var obj Customer
	var arrObj []Customer
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM customer"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Status)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetCustomerByUsername(username string) (Response, error) {
	var obj Customer
	var arrObj []Customer
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM customer WHERE customer.username LIKE '%" + username + "%'"
	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Status)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetCustomerByID(c_id int) (Customer, error) {

	var obj Customer

	con := db.CreateCon()
	fmt.Println(c_id)
	cid := strconv.Itoa(c_id)
	sqlStatement := "SELECT * FROM customer WHERE c_id = ?"

	rows := con.QueryRow(sqlStatement, cid)

	err := rows.Scan(&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Status)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func GetCustomerByID1(c_id int) (Response, error) {

	var obj Customer
	var arrObj []Customer
	var res Response

	con := db.CreateCon()
	fmt.Println(c_id)
	cid := strconv.Itoa(c_id)
	sqlStatement := "SELECT * FROM customer WHERE c_id = " + cid + ""

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Status)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

// insert data
func StoreCustomer(name string, username string, phone string, email string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO customer(name, username, phone, email, password, status) VALUES (?,?,?,?,?,'active')"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, username, phone, email, password)

	if err != nil {
		return res, err
	}

	//autoincrement (opsional)
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	//map[..] returnnya
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}

	return res, nil

}

// update
func UpdateCustomer(c_id int, name string, username string, phone string, email string, password string) (Response, error) {
	var res Response

	// p = controllers.GenerateHashPassword
	// password = p
	con := db.CreateCon()
	sqlStatement := "UPDATE customer SET name=?, username=?, phone=?, email=?, password=? WHERE customer.c_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, username, phone, email, password, c_id)

	if err != nil {
		return res, err
	}

	//autoincrement (opsional)
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update"
	//map[..] returnnya
	res.Data = map[string]int64{
		"row_affected": rowAffected,
	}

	return res, nil
}

// delete data
func DeleteCustomer(c_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE customer SET status = 'unactive' WHERE customer.c_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(c_id)

	if err != nil {
		return res, err
	}

	//autoincrement (opsional)
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Delete"
	//map[..] returnnya
	res.Data = map[string]int64{
		"row_affected": rowAffected,
	}

	return res, nil
}
