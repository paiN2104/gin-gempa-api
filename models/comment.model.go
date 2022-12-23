package models

import (
	"api_gempa/db"
	// "api_gempa/helpers"
	// "database/sql"
	// "github.com/go-playground/validator"
	"fmt"
	"net/http"
)

type Comment struct {
	Id        int       `json:"id"`
	Comment   string    `json:"comment"`
	userId	int       `json:"user_id"`
}

// func GetComments() ([]Comment, error) {
// 	var obj Comment
// 	var arrobj []Comment

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT * FROM comments"

// 	rows, err := con.Query(sqlStatement)
// 	return arrobj, err
// }

func addComment(comment string, userId int) (bool, error) {
	con := db.CreateCon()

	sqlStatement := "INSERT INTO comments (comment, user_id) VALUES (?, ?)"

	_, err := con.Exec(sqlStatement, comment, userId)

	if err != nil {
		fmt.Print("Query Error!")
		return false, err
	}

	return true, nil
}

func GetComments() (Response, error) {
	var obj Comment
	var arrobj []Comment
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM comments"

	rows, err := con.Query(sqlStatement)
	
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err := rows.Scan(&obj.Id, &obj.Comment, &obj.userId); err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

