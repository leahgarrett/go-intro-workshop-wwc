
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := 	"http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.81&lon=144.96&cnt=1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  fmt.Printf("Error reading weather body: %v", err)
	  return
	}
	fmt.Printf("Response: %s", body)
}
