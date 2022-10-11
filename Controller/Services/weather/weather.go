package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	kota "jwt-auth/Controller/Services/kota"
	"log"
	"net/http"
	"net/url"
)

type Cuaca struct{}

type Cordinate struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type EWeather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Entity struct {
	Name    string `json:"name"`
	Coord   Cordinate
	Weather []EWeather
}

func Init() *Cuaca {
	return &Cuaca{}
}

func (C *Cuaca) Weather(K *kota.KotaC) ([]byte, error) {
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

	var things Entity
	er := json.Unmarshal(body, &things)
	if er != nil {
		panic(er)
	}

	marshal, erro := json.Marshal(&things)
	if erro != nil {
		panic(er)
	}

	return marshal, err
}
