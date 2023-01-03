package models

import (
	"database/sql"
	"fmt"
	"vp_alp/db"
	"vp_alp/helpers"
)

// type User struct {
// 	Id       int    `json:"id"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }

// func CheckLoginTechnician1(username, password string) (bool, error) {
// 	var obj Technician
// 	var pwd string
// 	con := db.CreateCon()

// 	sqlStatement := "SELECT * FROM technician WHERE username = ? AND status='active'"
// 	err := con.QueryRow(sqlStatement, username).Scan(
// 		&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Rate, &obj.Kecamatan_id, &obj.Status, &obj.Kecamatan.K_id, &obj.Kecamatan.Kecamatan_name, &obj.Kecamatan.Wilayah_id)

// 	if err == sql.ErrNoRows {
// 		fmt.Print("Username not found!") //don't show in production env
// 		return false, err
// 	}

// 	if err != nil {
// 		fmt.Print("Query error!")
// 		return false, err
// 	}

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Print("Hash and password doesn't match!")
// 		return false, err
// 	}

// 	return true, nil
// }

func CheckLoginTechnician(username, password string) (int, error) {
	var obj Technician
	var pwd string
	var id int

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM technician WHERE username = ? AND status='active'"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Rate, &obj.Kecamatan_id, &obj.Status)

	// &t.T_id, &t.T_name, &t.Username, &t.Phone, &t.Email, &t.Password, &t.Rate, &t.Kecamatan_id, &t.Status, &t.Kecamatan.K_id, &t.Kecamatan.Kecamatan_name, &t.Kecamatan.Wilayah_id
	if err == sql.ErrNoRows {
		fmt.Print("Username not found!") //don't show in production env
		return 0, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return 0, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return 0, err
	}
	id = obj.T_id
	fmt.Println(id)

	return id, nil
}

func CheckLoginCustomer(username, password string) (int, error) {
	var obj Customer
	var pwd string
	var c_id int

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM customer WHERE username = ? AND status='active'"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Status)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!") //don't show in production env
		return 0, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return 0, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return 0, err
	}
	c_id = obj.C_id
	fmt.Println(c_id)

	return c_id, nil
}

// func CheckLoginCustomer(username, password string) (bool, error) {
// 	var obj Customer
// 	var pwd string
// 	con := db.CreateCon()

// 	sqlStatement := "SELECT * FROM customer WHERE username = ? AND status='active'"
// 	err := con.QueryRow(sqlStatement, username).Scan(
// 		&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Status,
// 	)

// 	if err == sql.ErrNoRows {
// 		fmt.Print("Username not found!") //don't show in production env
// 		return false, err
// 	}

// 	if err != nil {
// 		fmt.Print("Query error!")
// 		return false, err
// 	}

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Print("Hash and password doesn't match!")
// 		return false, err
// 	}

// 	return true, nil
// }
