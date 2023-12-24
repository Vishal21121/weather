package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		FeelsLike float64 `json:"feelslike_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		Humidity   float64 `json:"humidity"`
		AirQuality struct {
			PM25 float64 `json:"pm2_5"`
			PM10 float64 `json:"pm10"`
		} `json:"air_quality"`
	}
}

var generateCMD = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the weather for a given city",
	Long: `Fetches the weather for a given city
for Example: 
weather fetch -c London`,

	Run: func(cmd *cobra.Command, args []string) {
		city, _ := cmd.Flags().GetString("city")
		apiKey := os.Getenv("API_KEY")
		url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes", apiKey, city)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		data, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		var weatherData weather
		json.Unmarshal(data, &weatherData)
		fmt.Printf("The city is %s\n", weatherData.Location.Name)
		fmt.Printf("The Country is %s\n", weatherData.Location.Country)
		fmt.Printf("Temperature in celsius is %.2f\n", weatherData.Current.TempC)
		fmt.Printf("feels like in celsius is %.2f\n", weatherData.Current.FeelsLike)
		fmt.Printf("wind speed is %.2f Kmph\n", weatherData.Current.WindKph)
		fmt.Printf("Humidity is %.2f \n", weatherData.Current.Humidity)
		fmt.Printf("Condition of city is %s\n", weatherData.Current.Condition.Text)
		fmt.Printf("pm 2.5 data of city is %.2f\n", weatherData.Current.AirQuality.PM25)
		fmt.Printf("pm 10 data of city is %.2f\n", weatherData.Current.AirQuality.PM10)
	},
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rootCmd.AddCommand(generateCMD)
	generateCMD.Flags().StringP("city", "c", "London", "City for which to get the weather")
}
