package models

import (
    "database/sql"
    "fmt"
    "api_gempa/db"
    "api_gempa/helpers"
)

type User struct {
    Id	   int    `json:"id"`
    Name    string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func FetchLogin() ([]User, error) {
    var obj User
    var arrobj []User

    con := db.CreateCon()

    sqlStatement := "SELECT * FROM users"

    rows, err := con.Query(sqlStatement)

    if err != nil {
        return arrobj, err
    }

    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Password, &obj.Email)

        if err != nil {
            return arrobj, err
        }

        arrobj = append(arrobj, obj)
    }

    return arrobj, nil
}

func CheckLogin(username string, password string,email string) (bool, error) {
    var obj User
    var pwd string

    con := db.CreateCon()

    sqlStatement := "SELECT * FROM users WHERE username = ?"

    err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd, &obj.Email)

    if err == sql.ErrNoRows{
        fmt.Print("Username not found!") //dont show in production env
        return false, err
    }

    if err != nil {
        fmt.Print("Query Error!")
        return false, err
    }

    match, err := helpers.CheckPasswordHash(password, pwd)

    if !match{
        fmt.Print("Hash and Password not match!")
        return false, err
    }

    return true, nil
}

func Register(name string, username string, password string, email string,status string,image string) (bool, error) {
	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (name, username, password, email, status, image) VALUES (?, ?, ?, ?, ?, ?)"

	hash, _ := helpers.HashPassword(password)

	_, err := con.Exec(sqlStatement, name, username, hash, email, status, image)

	if err != nil {
		fmt.Print("Query Error!")
		return false, err
	}

	return true, nil
}