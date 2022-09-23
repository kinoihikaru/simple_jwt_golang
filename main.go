package main

import (
	auth "jwt-auth/Controller/Auth"
	db "jwt-auth/Db"
	model "jwt-auth/Model/User"
	route "jwt-auth/Route"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := db.Init().CreateConnection()

	model := model.Init(db)

	auth := auth.Init(model)

	var route = route.Init(auth)

	http.ListenAndServe(":3000", route.Run())

}
