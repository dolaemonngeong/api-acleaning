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

func CheckLoginTechnician(username, password string) (bool, error) {
	var obj Technician
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM technician WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.T_id, &obj.T_name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Rate, &obj.Kecamatan_id, &obj.Status,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!") //don't show in production env
		return false, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}

func CheckLoginCustomer(username, password string) (bool, error) {
	var obj Customer
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM customer WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.C_id, &obj.Name, &obj.Username, &obj.Phone, &obj.Email, &pwd, &obj.Status,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!") //don't show in production env
		return false, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}
