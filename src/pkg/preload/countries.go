package preload

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"math/rand"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/db"
	"os"
	"strconv"
	"time"
)

type Metadata struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

func preloadCountries() {
	var limit = config.GetMaxCountries()
	countries := getCountries()
	rand.Shuffle(len(*countries), func(i, j int) { (*countries)[i], (*countries)[j] = (*countries)[j], (*countries)[i] })
	for _, country := range (*countries)[:limit] {
		db.AddNewCountry(country)
	}
}

func RandomAddress() *db.Address {
	gofakeit.Seed(time.Now().UnixNano())
	address := gofakeit.Address()
	return &db.Address{
		Street:    address.Street,
		City:      address.City,
		State:     address.State,
		Zip:       address.Zip,
		Country:   address.Country,
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
	}
}

func getAddresses() []db.Address {
	var addresses []db.Address
	addressQuantity := rand.Intn(100)
	for i := 0; i < addressQuantity; i++ {
		address := RandomAddress()
		addresses = append(addresses, *address)
	}
	return addresses
}

func Pagination(page int) string {
	return "&page=" + strconv.Itoa(page)
}

func ParseResponse(country interface{}) *db.Country {
	countryMap := country.(map[string]interface{})
	latitude, _ := strconv.ParseFloat(countryMap["latitude"].(string), 64)
	longitude, _ := strconv.ParseFloat(countryMap["longitude"].(string), 64)
	regionMap := countryMap["region"].(map[string]interface{})
	region := db.Region{Id: regionMap["id"].(string), Name: regionMap["value"].(string)}
	addresses := getAddresses()
	warmingMetrics := InitializeWarmingMetrics()
	return &db.Country{Id: countryMap["id"].(string),
		Name:           countryMap["name"].(string),
		Capital:        countryMap["capitalCity"].(string),
		Latitude:       latitude,
		Longitude:      longitude,
		Region:         region,
		Addresses:      addresses,
		WarmingMetrics: warmingMetrics,
	}
}

func ParseMetadata(response interface{}) *Metadata {
	metadataString, _ := json.Marshal(&response)
	metadata := Metadata{}
	json.Unmarshal(metadataString, &metadata)
	return &metadata
}

func LoadCountriesByApi() *[]db.Country {
	var result []db.Country
	var partialResult []interface{}
	url := config.GetCountriesUrl()
	type CountryProviderResponse interface{}
	apiResponse := make([]CountryProviderResponse, 0)
	client := resty.New()
	for i := 1; i <= 6; i++ {
		client.R().EnableTrace().SetResult(&apiResponse).Get(url + Pagination(i))
		//metadata := ParseMetadata(apiResponse[0])
		partialResult = append(partialResult, apiResponse[1].([]interface{})...)
	}
	for _, country := range partialResult {
		result = append(result, *ParseResponse(country))
	}
	return &result
}

func CreateCountryFile(countries *[]db.Country, filepath string) {
	file, _ := json.MarshalIndent(countries, "", " ")
	_ = ioutil.WriteFile(filepath, file, 0644)
}

func LoadCountriesByFile(countries *[]db.Country, filepath string) *[]db.Country {
	file, _ := ioutil.ReadFile(filepath)
	_ = json.Unmarshal([]byte(file), &countries)
	return countries
}

func getCountries() *[]db.Country {
	var result *[]db.Country
	filepath := config.GetCountriesFile()

	if config.ShouldUpdateCountries() {
		result = LoadCountriesByApi()
		CreateCountryFile(result, filepath)
	} else {
		_, err := os.Stat(filepath)
		if os.IsNotExist(err) {
			fmt.Println("Countries do not exist in File")
			os.Exit(1)
		}
		result = LoadCountriesByFile(result, filepath)
	}

	//fmt.Println("Response Info:")
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", response.StatusCode())
	//fmt.Println("  Status     :", response.Status())
	//fmt.Println("  Proto      :", response.Proto())
	//fmt.Println("  Time       :", response.Time())
	//fmt.Println("  Received At:", response.ReceivedAt())
	//fmt.Println("  Body       :\n", response)

	return result
}
