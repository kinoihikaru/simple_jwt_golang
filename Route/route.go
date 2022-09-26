package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "jwt-auth/Controller/Auth"
	home "jwt-auth/Controller/Home"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Auth *auth.Auth
	Home *home.Home
}

func Init(a *auth.Auth, h *home.Home) *Route {
	return &Route{
		Auth: a,
		Home: h,
	}
}

func (R *Route) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t auth.LoginRequest
	decoder.Decode(&t)
	res, _ := R.Auth.Login(t)

	w.Write(res)
}

func (R *Route) home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	res := R.Home.Index()
	fmt.Println(res)
	w.Write(res)
}

func (R *Route) Run() *httprouter.Router {
	mux := httprouter.New()
	handler := auth.ParsingToken(R.home)
	mux.POST("/login", R.login)
	mux.GET("/home", handler)

	return mux
}
