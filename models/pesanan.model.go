package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"vp_alp/db"
)

type Pesanan struct {
	O_id       int            `json:"o_id" validate:"required"`
	Name       string         `json:"name" validate:"required"`
	Address    string         `json:"address" validate:"required"`
	Phone      string         `json:"phone" validate:"required"`
	Time       string         `json:"time" validate:"required"`
	Date       string         `json:"date" validate:"required"`
	Note       sql.NullString `json:"note"`
	Technician Technician
	T_id       uint   `json:"t_id" validate:"required"`
	C_id       int    `json:"c_id" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

// read all
func FetchAllOrder() (Response, error) {
	var obj Pesanan
	var arrObj []Pesanan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * from pesanan ORDER BY date DESC"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.O_id, &obj.Name, &obj.Address, &obj.Phone, &obj.Time, &obj.Date, &obj.Note, &obj.T_id, &obj.C_id, &obj.Status)

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
func StoreOrder(name string, address string, phone string, time string, date string, note string, t_id int, c_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO pesanan(name, address, phone, time, date, Note, t_id, c_id, status) VALUES (?,?,?,?,?,?,?,?,'Pending')"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phone, time, date, note, t_id, c_id)

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
func UpdateOrder(status string, o_id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE pesanan SET status=? WHERE pesanan.o_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(status, o_id)

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

func GetTechnicianOrder(t_id string, status string) (Response, error) {

	// var obj Technician
	var obj2 Pesanan
	var arrObj []Pesanan
	var res Response

	con := db.CreateCon()
	fmt.Println(status)
	fmt.Println(t_id)
	sqlStatement := "SELECT * FROM pesanan p INNER JOIN technician t ON t.t_id = p.t_id WHERE p.status = '" + status + "' AND p.t_id = '" + t_id + "'"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj2.O_id, &obj2.Name, &obj2.Address, &obj2.Phone, &obj2.Time, &obj2.Date, &obj2.Note, &obj2.T_id, &obj2.C_id, &obj2.Status, &obj2.Technician.T_id, &obj2.Technician.T_name, &obj2.Technician.Username, &obj2.Technician.Phone, &obj2.Technician.Email, &obj2.Technician.Password, &obj2.Technician.Rate, &obj2.Technician.Kecamatan_id, &obj2.Technician.Status)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj2)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetCustomerOrder(c_id string, status string) (Response, error) {

	// var obj Technician
	var obj2 Pesanan
	var arrObj []Pesanan
	var res Response

	con := db.CreateCon()
	fmt.Println(status)
	fmt.Println(c_id)
	sqlStatement := "SELECT * FROM pesanan p INNER JOIN technician t ON t.t_id = p.t_id WHERE p.status = '" + status + "' AND p.c_id = '" + c_id + "'"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj2.O_id, &obj2.Name, &obj2.Address, &obj2.Phone, &obj2.Time, &obj2.Date, &obj2.Note, &obj2.T_id, &obj2.C_id, &obj2.Status, &obj2.Technician.T_id, &obj2.Technician.T_name, &obj2.Technician.Username, &obj2.Technician.Phone, &obj2.Technician.Email, &obj2.Technician.Password, &obj2.Technician.Rate, &obj2.Technician.Kecamatan_id, &obj2.Technician.Status)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj2)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

// func GetCustomerOrder(c_id string, status string) (Response, error) {

// 	// var obj Customer
// 	var obj2 Pesanan
// 	var arrObj []Pesanan
// 	var res Response

// 	con := db.CreateCon()
// 	fmt.Println(status)
// 	fmt.Println(c_id)
// 	sqlStatement := "SELECT * FROM pesanan WHERE AND status = '" + status + "' AND c_id = '" + c_id + "'"

// 	rows, err := con.Query(sqlStatement)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj2.O_id, &obj2.Name, &obj2.Address, &obj2.Phone, &obj2.Time, &obj2.Date, &obj2.Note, &obj2.T_id, &obj2.C_id, &obj2.Status)

// 		if err != nil {
// 			return res, err
// 		}

// 		arrObj = append(arrObj, obj2)
// 	}

// 	res.Status = http.StatusOK
// 	res.Message = "Success"
// 	res.Data = arrObj

// 	return res, nil
// }

func GetOrderByID(o_id int) (Pesanan, error) {

	var obj Pesanan

	con := db.CreateCon()
	fmt.Println(o_id)
	oid := strconv.Itoa(o_id)
	sqlStatement := "SELECT * FROM pesanan WHERE o_id = ?"

	rows := con.QueryRow(sqlStatement, oid)

	err := rows.Scan(&obj.O_id, &obj.Name, &obj.Address, &obj.Phone, &obj.Time, &obj.Date, &obj.Note, &obj.T_id, &obj.C_id, &obj.Status)

	if err != nil {
		return obj, err
	}

	return obj, nil
}
