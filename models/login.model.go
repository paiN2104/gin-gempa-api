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
    Status string `json:"status"`
    Image string `json:"image"`
}

//get data by id
func FetchUserById(id string) (User, error) {
    var obj User

    con := db.CreateCon()

    sqlStatement := "SELECT * FROM users WHERE id = ?"

    rows := con.QueryRow(sqlStatement, id)

    err := rows.Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Password, &obj.Email, &obj.Status, &obj.Image)

    if err != nil {
        return obj, err
    }

    return obj, nil
}
    

//get data by id
// func FetchUserById(id string) (Response, error) {
// 	var obj User
// 	var arrObj []User
// 	var res Response

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT * FROM users WHERE id = ?"

// 	rows, err := con.Query(sqlStatement, id)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Password, &obj.Email, &obj.Status, &obj.Image)

// 		if err != nil {
// 			return res, err
// 		}

// 		arrObj = append(arrObj, obj)
// 	}

// 	res.Status = 200
// 	res.Message = "Success"
// 	res.Data = arrObj

// 	return res, nil
// }


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
        err = rows.Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Password, &obj.Email, &obj.Status, &obj.Image)

        if err != nil {
            return arrobj, err
        }

        arrobj = append(arrobj, obj)
    }

    return arrobj, nil
}

func CheckLogin(email string, password string) (int, error) {
    var obj User
    var pwd string
    var id int

    con := db.CreateCon()

    sqlStatement := "SELECT * FROM users WHERE email = ?"

    err := con.QueryRow(sqlStatement, email).Scan(&obj.Id, &obj.Name, &obj.Username, &pwd, &obj.Email, &obj.Status, &obj.Image)

    if err == sql.ErrNoRows{
        fmt.Print("Email not found!") //dont show in production env
        return 0, err
    }

    if err != nil {
        fmt.Print("Query Error!")
        return 0, err
    }

    match, err := helpers.CheckPasswordHash(password, pwd)

    if !match{
        fmt.Print("Hash and Password not match!")
        return 0, err
    }
    id = obj.Id

    return id, nil
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

//update user by id
func UpdateUser(id int, name string, username string, password string, email string, status string, image string) (bool, error) {
    con := db.CreateCon()

    sqlStatement := "UPDATE users SET name = ?, username = ?, password = ?, email = ?, status = ?, image = ? WHERE id = ?"

    _, err := con.Exec(sqlStatement, name, username, password, email, status, image, id)

    if err != nil {
        fmt.Print("Query Error!")
        return false, err
    }

    return true, nil
}