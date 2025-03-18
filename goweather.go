package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Main struct {
	Temp       float64 `json:"temp"`
	FeelsLike  float64 `json:"feels_like"`
	TempMin    float64 `json:"temp_min"`
	TempMax    float64 `json:"temp_max"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	SeaLevel   int     `json:"sea_level"`
	GrndLevel  int     `json:"grnd_level"`
}

type Sys struct {
	Type     int    `json:"type"`
	ID       int    `json:"id"`
	Country  string `json:"country"`
	Sunrise  int    `json:"sunrise"`
	Sunset   int    `json:"sunset"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherData struct {
	Coord      Coord       `json:"coord"`
	Weather    []Weather   `json:"weather"`
	Base       string      `json:"base"`
	Main       Main        `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       Wind        `json:"wind"`
	Clouds     Clouds      `json:"clouds"`
	DT         int         `json:"dt"`
	Sys        Sys         `json:"sys"`
	Timezone   int         `json:"timezone"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
}

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main2 struct {
	AQI int `json:"aqi"`
}

type Components struct {
	CO    float64 `json:"co"`
	NO    float64 `json:"no"`
	NO2   float64 `json:"no2"`
	O3    float64 `json:"o3"`
	SO2   float64 `json:"so2"`
	PM25  float64 `json:"pm2_5"`
	PM10  float64 `json:"pm10"`
	NH3   float64 `json:"nh3"`
}

type Item struct {
	Main      Main       `json:"main"`
	Components Components `json:"components"`
	Dt        int64      `json:"dt"`
}

type Response struct {
	Coord Coordinates `json:"coord"`
	List  []Item      `json:"list"`
}


func main() {

	var choice int
	var latitude, longitude string
	var city string

	content, err := os.ReadFile("api.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[0] : Entering latitude and longitude")
	fmt.Println("[1] : Entering city")
	fmt.Println("[Any other option] : Exit the program")

	fmt.Print("Enter choice: ")
	_, err = fmt.Scanf("%d", &choice)
	if err != nil {
		log.Fatalf("Error reading choice: %v", err)
	}

	if (choice == 0) {

		// Latitude
		fmt.Print("Enter latitude: ")
		_, err := fmt.Scanf("%v", &latitude)
		if err != nil {
			log.Fatalf("Error reading latitude: %v", err)
		}

		// Enter Longtitude
		fmt.Print("Enter longitude: ")
		_, err = fmt.Scanf("%v", &longitude)
		if err != nil {
			log.Fatalf("Error reading longitude: %v", err)
		}

		link := "https://api.openweathermap.org/data/2.5/weather?lat=" + string(latitude) + "&lon=" + string(longitude) + "&appid=" + string(content)
		
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var weatherData WeatherData

		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		fmt.Printf("Location: %s, %s\n", weatherData.Name, weatherData.Sys.Country)
		fmt.Printf("Weather: %s\n", weatherData.Weather[0].Description)
		fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp-273.15)
		fmt.Printf("Feels Like: %.2f°C\n", weatherData.Main.FeelsLike-273.15)
		fmt.Printf("Humidity: %d\n", weatherData.Main.Humidity)
		fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
		fmt.Printf("Pressure: %d hPa\n", weatherData.Main.Pressure)

		// Air pollution

		link2 := "http://api.openweathermap.org/data/2.5/air_pollution?lat=" + string(latitude) + "&lon=" + string(longitude) + "&appid=" + string(content)
		
		resp, err = http.Get(link2)
		if err != nil {
			log.Fatal(err)
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var components Components

		err = json.Unmarshal(body, &components)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}


		// fmt.Printf("Weather: %s\n", weatherData.Weather[0].Description)
		fmt.Printf("Сoncentration of CO: %.2f°C\n", Components.List[0].Components.CO)

	} else if (choice == 1) {

		// Enter City
		fmt.Print("Enter city: ")
		_, err = fmt.Scanf("%v", &city)
		if err != nil {
			log.Fatalf("Error reading city: %v", err)
		}

		link := "https://api.openweathermap.org/data/2.5/weather?q=" + string(city) + "&appid=" + string(content)
		
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var weatherData WeatherData

		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		fmt.Printf("Location: %s, %s\n", weatherData.Name, weatherData.Sys.Country)
		fmt.Printf("Weather: %s\n", weatherData.Weather[0].Description)
		fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp-273.15)
		fmt.Printf("Feels Like: %.2f°C\n", weatherData.Main.FeelsLike-273.15)
		fmt.Printf("Humidity: %d\n", weatherData.Main.Humidity)
		fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
		fmt.Printf("Pressure: %d hPa\n", weatherData.Main.Pressure)
	} else {
		fmt.Println("Invalid choice. Terminating the program")
		return
	}
}