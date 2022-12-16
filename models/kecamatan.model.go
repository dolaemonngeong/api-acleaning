package models

import (
	"net/http"
	"vp_alp/db"
)

type Kecamatan struct {
	K_id           int    `json:"k_id"`
	Kecamatan_name string `json:"kecamatan_name"`
	Wilayah_id     int    `json:"wilayah_id"`
}

// read all
func FetchAllKecamatan() (Response, error) {
	var obj Kecamatan
	var arrObj []Kecamatan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kecamatan"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.K_id, &obj.Kecamatan_name, &obj.Wilayah_id)

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
