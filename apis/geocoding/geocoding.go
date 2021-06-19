package geocoding

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const GeocodingEndpoint = "https://www.geocoding.jp/api"

func Geocoding(city string) (string, string) {
	values := url.Values{}
	values.Set("q", city)

	res, err := http.Get(fmt.Sprintf("%s?%s", GeocodingEndpoint, values.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var apiRes Geocode

	if err := xml.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
	return apiRes.Coordinate.Lat, apiRes.Coordinate.Lng
}

type Geocode struct {
	Address    string     `xml:"address"`
	Coordinate Coordinate `xml:"coordinate"`
}

type Coordinate struct {
	Lat string `xml:"lat"`
	Lng string `xml:"lng"`
}
