package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "jwt-auth/Controller/Auth"
	home "jwt-auth/Controller/Home"
	gmt "jwt-auth/Controller/Services/gmt"
	kota "jwt-auth/Controller/Services/kota"
	weather "jwt-auth/Controller/Services/weather"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Auth    *auth.Auth
	Home    *home.Home
	Weather *weather.Cuaca
	Gmt     *gmt.Gmt
}

func Init(a *auth.Auth, h *home.Home, w *weather.Cuaca, g *gmt.Gmt) *Route {
	return &Route{
		Auth:    a,
		Home:    h,
		Weather: w,
		Gmt:     g,
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

func (R *Route) weather(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	city := r.URL.Query().Get("city")
	var t kota.KotaC
	t.City = city
	var cuaca = weather.Cuaca{}
	resp, _ := cuaca.Weather(&t)

	w.Write(resp)
}

func (R *Route) gmt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	city := r.URL.Query().Get("city")
	var t kota.KotaC
	t.City = city
	var time = gmt.Gmt{}
	resp, _ := time.Timezone(&t)

	w.Write(resp)

}

func (R *Route) Run() *httprouter.Router {
	mux := httprouter.New()
	handler := auth.ParsingToken(R.home)
	mux.POST("/login", R.login)
	mux.GET("/home", handler)
	mux.GET("/cuaca", R.weather)
	mux.GET("/gmt", R.gmt)

	return mux
}
