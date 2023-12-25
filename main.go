package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
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
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				Chanceofrain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"Forecast"`
}

func main() {
	user := "London"

	if len(os.Args) >= 2 {
		user = os.Args[1]
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=dda74783104147c4a99175136232512&q=" + user + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather Api not working")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	var weather Weather

	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}
	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf("Currently %s %s %0.fC %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		// println(hour.TimeEpoch)

		if date.Before(time.Now()) {
			continue
		}
		messaage := fmt.Sprintf("%s %0.fC %0.f%% %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.Chanceofrain,
			hour.Condition.Text,
		)
		if hour.Chanceofrain > 40 {
			color.Red(messaage)
		} else {
			color.Green(messaage)
		}
	}

}
