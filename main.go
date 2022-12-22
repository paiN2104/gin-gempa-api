package main

import (
	"api_gempa/routes"
	"api_gempa/db"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}