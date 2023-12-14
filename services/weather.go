package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"weather-cli/config"
	"weather-cli/models"
	"weather-cli/utils"
)

type WeatherProvider interface {
	GetWeatherForecast(location string) (weather models.Weather)
}

type WeatherServices struct {
	Config *config.AppConfig
}

var (
	once     sync.Once
	instance *WeatherServices
)

func (w *WeatherServices) GetWeatherForecast(location string) (weather models.Weather) {
	var API_KEY string = w.Config.WeatherAPIKey

	filename := fmt.Sprintf("%s_%s.json", strings.ToLower(location), time.Now().Format("2006-01-02"))

	if utils.FileExistsT(filename) {
		data, err := utils.ReadFileT(filename)

		if err != nil {
			return models.Weather{}
		}

		err = json.Unmarshal(data, &weather)
		if err != nil {
			return models.Weather{}
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

func NewWeatherServices(config *config.AppConfig) *WeatherServices {
	once.Do(func() {
		instance = &WeatherServices{
			Config: config,
		}
	})

	return instance
}

func GetInstance() *WeatherServices {
	return instance
}
