This is a Go CLI program that fetches weather data via OpenWeatherMap API.

Create an api.txt file containing the API Key from https://home.openweathermap.org/api_keys

According to (https://openweathermap.org/faq), "Your API key will be activated automatically, up to 2 hours after your successful registration."

```shell
git clone https://github.com/mariosdaskalas/goweather && cd goweather && touch api.txt && echo -n "PLACE_YOUR_API_KEY_HERE" > api.txt && go run .
```

```shell
go version go1.24.1 linux/amd64
```

Use https://www.latlong.net to find Latitude and Longitude of a Location