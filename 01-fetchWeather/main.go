package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "57d8d796ede7bbba36296f950b1fe02f"

var data struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func main() {
	start := time.Now()
	cities := []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix",
		"Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose",
		"Austin", "Jacksonville", "Fort Worth", "Columbus", "San Francisco",
		"Charlotte", "Indianapolis", "Seattle", "Denver", "Washington",
		"Boston", "El Paso", "Detroit", "Nashville", "Portland",
		"Memphis", "Oklahoma City", "Las Vegas", "Louisville", "Baltimore",
		"Milwaukee", "Albuquerque", "Tucson", "Fresno", "Sacramento",
		"Mesa", "Kansas City", "Atlanta", "Long Beach", "Colorado Springs",
		"Miami", "Raleigh", "Omaha", "Minneapolis", "Tulsa",
		"Cleveland", "Wichita", "Arlington", "New Orleans", "Bakersfield",
		"Tampa", "Honolulu", "Anaheim", "Aurora", "Santa Ana",
		"St. Louis", "Riverside", "Corpus Christi", "Lexington", "Pittsburgh",
		"Anchorage", "Stockton", "Cincinnati", "St. Paul", "Toledo",
		"Greensboro", "Newark", "Plano", "Henderson", "Lincoln",
		"Orlando", "Jersey City", "Chula Vista", "Buffalo", "Fort Wayne",
		"Chandler", "St. Petersburg", "Laredo", "Durham", "Irvine",
		"Madison", "Norfolk", "Lubbock", "Gilbert", "Winston-Salem",
		"Glendale", "Reno", "Hialeah", "Garland", "Chesapeake",
		"Irving", "North Las Vegas", "Scottsdale", "Baton Rouge", "Fremont"}

	ch := make(chan string)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(start))
}
