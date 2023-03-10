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
	UserId    int       `json:"user_id"`
	Name	string		`json:"name"`
}

// func GetComments() ([]Comment, error) {
// 	var obj Comment
// 	var arrobj []Comment

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT * FROM comments"

// 	rows, err := con.Query(sqlStatement)
// 	return arrobj, err
// }

func StoreComment(comment string, userId string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO comments (comment, user_id) VALUES (?, ?)"

	_, err := con.Exec(sqlStatement, comment, userId)

	if err != nil {
		fmt.Print("Query Error!")
		return res, err
	}

	return res, nil
}

func GetAllComments() (Response, error) {
	var obj Comment
	// var objUser User
	var arrobj []Comment
	// var arrobjUser []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT a.name,b.comment,b.id,b.user_id FROM users a, comments b WHERE a.id = b.user_id ORDER BY b.id DESC LIMIT 10"

	rows, err := con.Query(sqlStatement)
	
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err := rows.Scan(&obj.Name, &obj.Comment ,&obj.Id, &obj.UserId); err != nil {
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

func DeleteComment(id string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM comments WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)

	if (err != nil) {
		return res, err
	}
	rowAffected, err:= result.RowsAffected()

	if (err != nil) {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowAffected,
	}

	return res, nil
}

