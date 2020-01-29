
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":5000", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	count := r.FormValue("count")

	body, err := getWeatherResponseBody(count)

	if err != nil {
		panic(err)
	}

	openWeather := OpenWeather{}
	err = json.Unmarshal(body, &openWeather)
	if err != nil {
		panic(err)
	}
	
	for i := range openWeather.List  {
		fmt.Fprintf(w, "\nWeather in %s is %.2f", 
		openWeather.List[i].Name,
		openWeather.List[i].Temperature.NormalisedCurrentTemp())
	}
}
	

func (t TemperatureDetails) NormalisedCurrentTemp() float64 {
	return t.CurrentTemp - 273.15
}


type OpenWeather struct {
	List []City `json:"list"`
}

type City struct {
	Temperature TemperatureDetails `json:"main"`
	Name    string  `json:"name"`
}

type TemperatureDetails struct {
	CurrentTemp float64 `json:"temp"`
	MaxTemp     float64 `json:"temp_max"`
}

func getWeatherResponseBody(count string) ([]byte, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.81&lon=144.96&cnt=%s", count)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return []byte(""), err
	}
defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
     
	if err != nil {
		fmt.Printf("Error reading weather: %v", err)
		return []byte(""), err
	}
	
	return body, nil
}
