package gmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	Kota "jwt-auth/Controller/Services/kota"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Gmt struct{}

type RespG struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Timezone int64  `json:"timezone"`
	Cod      int    `json:"cod"`
}

type EntityG struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Timezone time.Time `json:"timezone"`
	Cod      int       `json:"cod"`
}

func Init() *Gmt {
	return &Gmt{}
}
func (g *Gmt) Timezone(K *Kota.KotaC) ([]byte, error) {
	client := &http.Client{}

	city := K.City
	api_key := "25098951bbc63c36ddce74381b182c14"

	params := "q=" + url.QueryEscape(city) + "&appid=" + url.QueryEscape(api_key)
	path := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?%s", params)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var things RespG
	er := json.Unmarshal(body, &things)
	if er != nil {
		panic(er)
	}

	tm := time.Unix(things.Timezone, 0)

	var convert = EntityG{
		Id:       things.Id,
		Name:     things.Name,
		Timezone: tm,
		Cod:      things.Cod,
	}

	var data, _ = json.Marshal(&convert)

	return data, err
}
