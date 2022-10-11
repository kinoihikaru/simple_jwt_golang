package timezone

import (
	"fmt"
	"io/ioutil"
	Kota "jwt-auth/Controller/Services/kota"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type TimeZone struct{}

func (t *TimeZone) Gmt(c Kota.ApiCity) ([]byte, error) {
	client := &http.Client{}

	var lat string = strconv.FormatFloat(float64(c.Latitude), 'f', 12, 64)
	var lng string = strconv.FormatFloat(float64(c.Longitude), 'f', 12, 64)

	api_key := "5E6BYBRMSI2I"

	params := "key=" + url.QueryEscape(api_key) + "&format=json" + "&lat=" + url.QueryEscape(lat) + "&lng=" + url.QueryEscape(lng)
	path := fmt.Sprintf("http://api.timezonedb.com/v2.1/get-time-zone?%s", params)

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
