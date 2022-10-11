package main

import (
	auth "jwt-auth/Controller/Auth"
	home "jwt-auth/Controller/Home"
	gmt "jwt-auth/Controller/Services/gmt"
	weather "jwt-auth/Controller/Services/weather"
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
	home := home.Init()

	weather := weather.Init()

	gmt := gmt.Init()

	var route = route.Init(auth, home, weather, gmt)

	http.ListenAndServe(":3000", route.Run())

}
