package job

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"org.gene/singularity/pkg/config"
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
