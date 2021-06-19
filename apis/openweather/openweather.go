package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ta93-ito/discord-weather-bot/apis/geocoding"
	"github.com/ta93-ito/discord-weather-bot/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Endpoint = "https://api.openweathermap.org/data/2.5/forecast"

func GetForecast(city string) (ForecastList, error) {
	lat, lon := geocoding.Geocoding(city)

	token := config.Config.ApiKey

	values := url.Values{}
	values.Set("lat", lat)
	values.Set("lon", lon)
	values.Set("appid", token)
	values.Set("cnt", "7")
	values.Set("lang", "ja")

	res, err := http.Get(fmt.Sprintf("%s?%s", Endpoint, values.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	switch res.Status[0:1] {
	case "4":
		var nilRes ForecastList
		return nilRes, errors.New("invalid statement!")
	case "5":
		var nilRes ForecastList
		return nilRes, errors.New("something went wrong...")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var apiRes ForecastList

	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		fmt.Println(err)
	}
	var nonErr error
	return apiRes, nonErr
}

type ForecastList struct {
	Forecasts []Forecast `json:"list"`
}

type Forecast struct {
	Main    Main    `json:"main"`
	Weather Weather `json:"weather"`
	DtTxt   string  `json:"dt_txt"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type Weather []struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
