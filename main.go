package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
				TimeEpoch int64   `json:"timeepoch"`
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
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=dda74783104147c4a99175136232512&q=Lucknow&days=1&aqi=no&alerts=no")

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
	// fmt.Println(weather)
	location, current, _ := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf("Currently %s %s %0.fC %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

}
