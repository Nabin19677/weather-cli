/*
Copyright © 2023 Anil Khadka
*/
package cmd

import (
	"fmt"
	"time"
	"weather-cli/services"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Location string

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: `View Today's Weather Forecast.`,
	Run: func(cmd *cobra.Command, args []string) {
		weatherService := &services.WeatherServices{}
		weather := weatherService.GetWeatherForecast(Location)
		location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

		color.Yellow("%s, %s: %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)

		for _, hour := range hours {
			date := time.Unix(hour.TimeEpoch, 0)

			if date.Before(time.Now()) {
				continue
			}

			message := fmt.Sprintf("%s - %0.fC, %0.f%%, %s\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

			if hour.ChanceOfRain >= 40 {
				color.Red(message)
			} else {
				color.Blue(message)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().StringVarP(&Location, "location", "l", "Kathmandu", "Location to view weather forecast")
}
