package suhu_cuaca

import (
	"fmt"
	"io/ioutil"
	Kota "jwt-auth/Controller/Services/kota"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Weather struct{}

func (W *Weather) Cuaca(c Kota.ApiCity) ([]byte, error) {
	client := &http.Client{}

	var lat string = strconv.FormatFloat(float64(c.Latitude), 'f', 12, 64)
	var long string = strconv.FormatFloat(float64(c.Longitude), 'f', 12, 64)

	api_key := "25098951bbc63c36ddce74381b182c14"

	params := "lat=" + url.QueryEscape(lat) + "&lon=" + url.QueryEscape(long) + "&appid=" + url.QueryEscape(api_key)
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

	return body, err
}
