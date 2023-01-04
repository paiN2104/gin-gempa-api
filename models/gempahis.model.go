package models

import (
	"api_gempa/db"
	"fmt"
	"net/http"
)

type Gempa struct {
	Id        int       `json:"id"`
	Wilayah   string    `json:"wilayah"`
	Tanggal   string    `json:"tanggal"`
	Magnitudo string    `json:"magnitudo"`
}
//store data
func StoreGempaHis(Wilayah string, Tanggal string, Magnitudo string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO gempahis (wilayah, tanggal, magnitudo) VALUES (?, ?, ?)"

	_, err := con.Exec(sqlStatement, Wilayah, Tanggal, Magnitudo)

	if err != nil {
		fmt.Print("Query Error!")
		return res, err
	}

	return res, nil
}

func GetGempaHis()(Response, error){
	var obj Gempa
	var arrobj []Gempa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM gempahis"

	rows, err := con.Query(sqlStatement)
	
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err := rows.Scan(&obj.Id,&obj.Magnitudo, &obj.Tanggal ,&obj.Wilayah); err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
		// arrobjUser = append(arrobjUser)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil

}