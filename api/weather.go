package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"weather-cli/utils"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func GetWeatherForecast(location string) (weather Weather) {
	var API_KEY string = os.Getenv("WEATHER_API_KEY")

	filename := fmt.Sprintf("%s_%s.json", location, time.Now().Format("2006-01-02"))

	if utils.FileExistsT(filename) {
		fmt.Println("FIle Exists")
		data, err := os.ReadFile(filename)

		if err != nil {
			return Weather{}
		}

		err = json.Unmarshal(data, &weather)
		if err != nil {
			return Weather{}
		}

		return

	} else {
		res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=" + location + "&days=1&aqi=no&alerts=no")

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			panic("Weather API is not available")
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(body, &weather)
		if err != nil {
			panic(err)
		}

		err = utils.WriteToFileT(filename, body)

		if err != nil {
			panic(err)
		}

		return
	}
}
