package models

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "vp_alp/controllers"
	"vp_alp/db"
)

type Technician struct {
	T_id         int    `json:"t_id" validate:"required"`
	T_name       string `json:"t_name" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Rate         int    `json:"rate" validate:"required"`
	Kecamatan_id int    `json:"kecamatan_id" validate:"required"`
	Status       string `json:"status" validate:"required"`
}

// read all
func FetchAllTechnician() (Response, error) {
	var obj Technician
	var arrObj []Technician
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM technician"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Rate, &obj.Kecamatan_id, &obj.Status)

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

func GetTechnicianByName(t_name string) (Response, error) {

	//sqlStatement := "SELECT name FROM technician WHERE name LIKE '%Songa%'"

	var obj Technician
	var arrObj []Technician
	var res Response

	con := db.CreateCon()
	fmt.Println(t_name)
	sqlStatement := "SELECT * FROM technician WHERE t_name LIKE '%" + t_name + "%'"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Rate, &obj.Kecamatan_id, &obj.Status)

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

func GetTechnicianByLocation(kid string) (Response, error) {

	var obj Technician
	var obj2 Kecamatan
	var arrObj []Technician
	var res Response

	con := db.CreateCon()
	// kid := strconv.Itoa(k_id)
	fmt.Println(kid)
	sqlStatement := "SELECT * FROM technician t, kecamatan k WHERE kecamatan_name = '" + kid + "' AND t.kecamatan_id = k.kecamatan_id "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Rate, &obj.Kecamatan_id, &obj2.K_id, &obj2.Kecamatan_name, &obj2.Wilayah_id)

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

func GetTechnicianByID(t_id int) (Response, error) {

	var obj Technician
	var arrObj []Technician
	var res Response

	con := db.CreateCon()
	fmt.Println(t_id)
	tid := strconv.Itoa(t_id)
	sqlStatement := "SELECT * FROM technician WHERE t_id = " + tid + ""

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &obj.Password, &obj.Rate, &obj.Kecamatan_id)

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
func StoreTechnician(t_name string, username string, phone string, email string, password string, kecamatan_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO technician(t_name, username, phone, email, password, kecamatan_id, status) VALUES (?,?,?,?,?,?,'active')"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(t_name, username, phone, email, password, kecamatan_id)

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
func UpdateTechnician(t_id int, t_name string, username string, phone string, email string, password string, kecamatan_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE technician SET t_name=?, username=?, phone=?, email=?, password=?, kecamatan_id=? WHERE technician.t_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(t_name, username, phone, email, password, kecamatan_id, t_id)

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

func UpdateTechnicianRate(rate int, t_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE technician SET rate=? WHERE technician.t_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(rate, t_id)

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
func DeleteTechnician(t_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE technician SET status = 'unactive' WHERE technician.t_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(t_id)

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
