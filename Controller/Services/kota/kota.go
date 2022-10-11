package kota

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type City struct {
}

type KotaC struct {
	City string `json:"city"`
}

type ApiCity struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type PositionCity struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (C *City) KotaController(k *KotaC) ApiCity {
	client := &http.Client{}
	name := k.City

	params := "city=" + url.QueryEscape(name)
	path := fmt.Sprintf("https://api.api-ninjas.com/v1/geocoding?%s", params)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Api-Key", "BfIVcKVCNQb0CpJvIfyHeg==EpKWVZ5f1GDOqxVJ")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var things []PositionCity
	er := json.Unmarshal(body, &things)
	if er != nil {
		panic(er)
	}

	var data = ApiCity{
		Latitude:  things[0].Latitude,
		Longitude: things[0].Longitude,
	}

	return data
}
