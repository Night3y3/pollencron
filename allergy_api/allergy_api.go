package allergy_api

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/Night3y3/pollencron/utils"
	"github.com/joho/godotenv"
)

type HourlyLoadResult struct {
	Total  int   `json:"total"`
	Hourly []int `json:"hourly"`
}

type HourlyLoadResponse struct {
	Success int              `json:"success"`
	Result  HourlyLoadResult `json:"result"`
}

type CurrentChartDataResult struct {
	Date    string  `json:"date"`
	Average float64 `json:"average"`
}

type CurrentChartDataResponse struct {
	Success int                      `json:"success"`
	Results []CurrentChartDataResult `json:"results"`
}

func GetHourlyLoadData() (*string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil,err
	}

	api_url := os.Getenv("API_URL")

	hourly_url := api_url+"/api/v1/pollen/hourly"
	response, err := utils.MakeHTTPRequest(hourly_url, "GET", nil, nil, nil,HourlyLoadResponse{})
	if err != nil {
		return nil, err
	}
	averageLoad := 0
	for _, hourlyLoad := range response.Result.Hourly {
		averageLoad += hourlyLoad
	}
	averageLoad = averageLoad / len(response.Result.Hourly)

	scaledAverageLoad := averageLoad/2

	formattedMessage := fmt.Sprintf("The average pollen load for today is %d, which is %d%% of the maximum load.", averageLoad, scaledAverageLoad)

	return &formattedMessage, nil
}
 
func GetCurrentChartData() (*string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil,err
	}

	api_url := os.Getenv("API_URL")

	chart_url := api_url+"/api/v1/pollen/chart"

	response, err := utils.MakeHTTPRequest(chart_url, "GET", nil, nil, nil,CurrentChartDataResponse{})
	if err != nil {
		return nil, err
	}

	currentYYYYMMDD := time.Now().Format("2006-01-02")
	averageHistorical := 0.0
	for _, currentChartData := range response.Results {
		if(currentChartData.Date == currentYYYYMMDD) {
			averageHistorical = currentChartData.Average
		}
	}

	scaledAverageHistorical := int(math.Round(averageHistorical/2.0))
	
	formattedMessage := fmt.Sprintf("The average pollen load for today is %f, which is %d%% of the maximum load.", averageHistorical, scaledAverageHistorical)

	return &formattedMessage, nil
}