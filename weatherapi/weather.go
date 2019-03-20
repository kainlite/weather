package weatherapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Location struct {
	LocationId string
	UserId     string
}

// We are doing this while learning about the API
// to save some requests
func readJSONFile(name string, forecast interface{}) {
	// Open our jsonFile
	jsonFile, err := os.Open(name)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened File")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &forecast)

	fmt.Printf("%+v\n", forecast)
}

func PrintStats() {
	var hourly HourlyForecast
	var daily DailyForecast

	readJSONFile("example_by_hour.json", &hourly)
	readJSONFile("example_by_day.json", &daily)

	// Hourly
	// A few notes here, notice the 20 + v.Num (20 was the hour of the query,
	// then Num will be a sequential number from 1 to 48 which are the following 48 hours)
	// on the other side I forgot to set the current unit type and the default is Fahrenheit
	// so we convert it with the formula (F - 32) * 5/9.
	for _, v := range hourly.Forecasts {
		fmt.Printf("Day: %s, Hour: %d, Forecast: %s - %d\n", v.Dow, ((20 + v.Num) % 24), v.Phrase22Char, ((v.Temp - 32) * 5 / 9))
	}

	// Daily
	for _, v := range daily.Forecasts {
		fmt.Printf("Day: %s, Forecast: %s, Min: %d, Max: %f\n", v.Dow, v.Narrative, v.MinTemp, v.MaxTemp)
	}
}

// TODO
func CreateLocation(LocationId string, UserId string) Location {
	return Location{}
}

func DeleteLocation(LocationId string) Location {
	return Location{}
}

func SearchLocation(Name string) Location {
	return Location{}
}

func Dashboard() []Location {
	return []Location{}
}
