package job

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/margostino/singularity/pkg/config"
	"github.com/margostino/singularity/pkg/db"
)

type AirQualityResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

func GetAirQualityFor(latitude float64, longitude float64) float64 {
	var apiResponse AirQualityResponse
	config := config.GetJobsConfigurationBy("air_quality")
	lat := fmt.Sprintf("%f", latitude)
	long := fmt.Sprintf("%f", longitude)
	url := fmt.Sprintf("%s/geo:%s;%s/?token=%s", config.Url, lat, long, config.Token)
	client := resty.New()
	client.R().EnableTrace().SetResult(&apiResponse).Get(url)
	// TODO: validate apiResponse.Status = "error"
	return apiResponse.Data["aqi"].(float64)
}

func UpdateAirQuality() {
	for _, country := range *db.GetCountries() {
		for index := range country.WarmingMetrics {
			if country.WarmingMetrics[index].Key == "air_quality" {
				value := GetAirQualityFor(country.Latitude, country.Longitude)
				country.WarmingMetrics[index].Value = value
			}
		}
	}
}
