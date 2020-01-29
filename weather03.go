package main

import (
	"fmt"
)

func main() {
	url := "http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.81&lon=144.96&cnt=1"

	fmt.Printf("Our API url is %s", url)
}
