// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// var name string
// 	// fmt.Println("Enter your name: ")
// 	// fmt.Scanln(&name)
// 	// fmt.Println("Your name is", name)

// 	wellcome := "Wellcome to user input"
// 	fmt.Println(wellcome)

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter your full name: ")

// 	//comma ok || err ok

// 	// fullname, _ := reader.ReadString('\n')
// 	// fmt.Println("Your full name is", fullname)
// 	// fmt.Printf("type of name is %T", fullname)

// 	fullname, _, _ := reader.ReadLine()
// 	fmt.Println("Your full name is", string(fullname))
// 	fmt.Printf("type of name is %T", fullname)

// }
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CityConditionResponse struct {
	City1     string `json:"City1"`
	City2     string `json:"City2"`
	Condition string `json:"Condition"`
}

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Wind        float64 `json:"wind"`
	Rain        float64 `json:"rain"`
	Cloud       float64 `json:"cloud"`
}

func getCityAndCondition() (*CityConditionResponse, error) {
	url := "https://quest.squadcast.tech/api/ENG21CS0030/weather"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cityCondition CityConditionResponse
	if err := json.NewDecoder(resp.Body).Decode(&cityCondition); err != nil {
		return nil, err
	}

	return &cityCondition, nil
}

func getWeatherDetails(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://quest.squadcast.tech/api/ENG21CS0030/weather/get?q=%s", city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func determineBetterCity(city1, city2 string, weather1, weather2 *WeatherResponse, condition string) (string, error) {
	switch condition {
	case "hot":
		if weather1.Temperature > weather2.Temperature {
			return city1, nil
		} else {
			return city2, nil
		}
	case "cold":
		if weather1.Temperature < weather2.Temperature {
			return city1, nil
		} else {
			return city2, nil
		}
	case "windy":
		if weather1.Wind > weather2.Wind {
			return city1, nil
		} else {
			return city2, nil
		}
	case "rainy":
		if weather1.Rain > weather2.Rain {
			return city1, nil
		} else {
			return city2, nil
		}
	case "sunny":
		if weather1.Cloud < weather2.Cloud {
			return city1, nil
		} else {
			return city2, nil
		}
	case "cloudy":
		if weather1.Cloud > weather2.Cloud {
			return city1, nil
		} else {
			return city2, nil
		}
	default:
		return "", fmt.Errorf("invalid condition: %s", condition)
	}
}

func main() {
	cityCondition, err := getCityAndCondition()
	if err != nil {
		log.Fatalf("Error fetching city and condition: %v", err)
	}

	weather1, err := getWeatherDetails(cityCondition.City1)
	if err != nil {
		log.Fatalf("Error fetching weather data for %s: %v", cityCondition.City1, err)
	}

	weather2, err := getWeatherDetails(cityCondition.City2)
	if err != nil {
		log.Fatalf("Error fetching weather data for %s: %v", cityCondition.City2, err)
	}

	betterCity, err := determineBetterCity(cityCondition.City1, cityCondition.City2, weather1, weather2, cityCondition.Condition)
	if err != nil {
		log.Fatalf("Error determining better city: %v", err)
	}

	fmt.Printf("The better city based on the condition '%s' is: %s\n", cityCondition.Condition, betterCity)
}
