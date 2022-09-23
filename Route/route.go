package route

import (
	"encoding/json"
	"net/http"

	auth "jwt-auth/Controller/Auth"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Auth *auth.Auth
}

func Init(a *auth.Auth) *Route {
	return &Route{
		Auth: a,
	}
}

func (R *Route) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var res ,err = R.Auth.Login()
	decoder := json.NewDecoder(r.Body)
	var t auth.LoginRequest
	decoder.Decode(&t)
	res, _ := R.Auth.Login(t)

	w.Write(res)
}

func (R *Route) Run() *httprouter.Router {
	mux := httprouter.New()

	mux.POST("/login", R.login)

	return mux
}
