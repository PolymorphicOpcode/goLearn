package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/jedib0t/go-pretty/v6/table"
)

type WeatherProperties struct {
	Periods []struct {
		StartTime       string  `json:"startTime"`
		Temperature     float64 `json:"temperature"`
		TemperatureUnit string  `json:"temperatureUnit"`
		ShortForecast   string  `json:"shortForecast"`
	} `json:"periods"`
}

type Weather struct {
	Properties WeatherProperties `json:"properties"`
}

type PointResponse struct {
	Properties struct {
		Forecast         string `json:"forecast"`
		RelativeLocation struct {
			Properties struct {
				City  string `json:"city"`
				State string `json:"state"`
			} `json:"properties"`
		} `json:"relativeLocation"`
	} `json:"properties"`
}

var (
	contxt  = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)

func main() {
	location := "33.7490,-84.3880" // Default coordinates for Atlanta
	if len(os.Args) >= 2 {
		location = strings.Join(os.Args[1:], " ")
	}
	locationEncoded := url.QueryEscape(location)
	key := fmt.Sprintf("weatherData:%s", locationEncoded)

	val, err := RedisDb.Get(contxt, key).Result()
	if err == redis.Nil {
		val = fetchWeatherData(location)
		if err := RedisDb.Set(contxt, key, val, 10*time.Minute).Err(); err != nil {
			log.Fatalf("Failed to set data in Redis: %v", err)
		}
	} else if err != nil {
		log.Fatalf("Failed to get data from Redis: %v", err)
	}

	var weather Weather
	if err := json.Unmarshal([]byte(val), &weather); err != nil {
		log.Fatalf("Failed to unmarshal weather data: %v", err)
	}

	displayWeather(weather)
}

func fetchWeatherData(location string) string {
	pointURL := fmt.Sprintf("https://api.weather.gov/points/%s", location)
	resp, err := http.Get(pointURL)
	if err != nil {
		log.Fatalf("Failed to fetch point data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Weather API not available")
	}

	var pointData PointResponse
	if err := json.NewDecoder(resp.Body).Decode(&pointData); err != nil {
		log.Fatalf("Failed to decode point data: %v", err)
	}

	forecastURL := pointData.Properties.Forecast

	// Fetch weather forecast data from the forecast URL
	resp, err = http.Get(forecastURL)
	if err != nil {
		log.Fatalf("Failed to fetch forecast data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Weather API not available")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return string(body)
}

func displayWeather(weather Weather) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleBold)
	t.Style().Box = table.StyleBoxLight
	t.Style().Box.Left = " "
	t.Style().Box.Right = " "

	t.AppendHeader(table.Row{"Time", "Temp (°F)", "Condition"})

	for _, period := range weather.Properties.Periods {
		startTime, err := time.Parse(time.RFC3339, period.StartTime)
		if err != nil {
			log.Fatalf("Failed to parse time: %v", err)
		}

		tempColor := getColor(period.Temperature)
		t.AppendRow(table.Row{
			startTime.Format("2006-01-02 15:04"),
			tempColor.Sprintf("%.0f°F", period.Temperature),
			getWeatherEmoji(period.ShortForecast),
		})
	}

	t.Render()
}

func getColor(temp float64) *color.Color {
	switch {
	case temp <= 32:
		return color.New(color.FgBlue)
	case temp > 32 && temp <= 50:
		return color.New(color.FgCyan)
	case temp > 50 && temp <= 68:
		return color.New(color.FgGreen)
	case temp > 68 && temp <= 86:
		return color.New(color.FgYellow)
	case temp > 86:
		return color.New(color.FgRed)
	default:
		return color.New(color.FgWhite)
	}
}

func getWeatherEmoji(description string) string {
	description = strings.ToLower(description)
	switch {
	case strings.Contains(description, "sunny"):
		return "☀️"
	case strings.Contains(description, "clear"):
		return "🌤️"
	case strings.Contains(description, "cloudy"):
		return "☁️"
	case strings.Contains(description, "rain"):
		return "🌧️"
	case strings.Contains(description, "showers"):
		return "🌦️"
	case strings.Contains(description, "thunderstorm"):
		return "⛈️"
	case strings.Contains(description, "snow"):
		return "❄️"
	case strings.Contains(description, "fog"):
		return "🌫️"
	default:
		return "❓"
	}
}
