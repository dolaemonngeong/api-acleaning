package models

import (
	"net/http"
	"vp_alp/db"
)

type Wilayah struct {
	W_id int    `json:"w_id"`
	Name string `json:"name"`
}

// read all
func FetchAllWilayah() (Response, error) {
	var obj Wilayah
	var arrObj []Wilayah
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM wilayah"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.W_id, &obj.Name)

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
