package preload

import (
	"encoding/json"
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
	countries := getCountries()
	for _, country := range countries {
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
	regionMap := countryMap["region"].(map[string]interface{})
	region := db.Region{Id: regionMap["id"].(string), Name: regionMap["value"].(string)}
	latitude, _ := strconv.ParseFloat(countryMap["latitude"].(string), 64)
	longitude, _ := strconv.ParseFloat(countryMap["longitude"].(string), 64)
	addresses := getAddresses()
	return &db.Country{Id: countryMap["id"].(string),
		Name:      countryMap["name"].(string),
		Capital:   countryMap["capitalCity"].(string),
		Latitude:  latitude,
		Longitude: longitude,
		Region:    region,
		Addresses: addresses,
	}
}

func ParseMetadata(response interface{}) *Metadata {
	metadataString, _ := json.Marshal(&response)
	metadata := Metadata{}
	json.Unmarshal(metadataString, &metadata)
	return &metadata
}

func getCountries() []db.Country {
	var result []db.Country
	filepath := config.GetCountriesFile()
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		url := config.GetCountriesUrl()
		type CountryProviderResponse interface{}
		apiResponse := make([]CountryProviderResponse, 0)
		client := resty.New()
		client.R().EnableTrace().SetResult(&apiResponse).Get(url + Pagination(1))
		metadata := ParseMetadata(apiResponse[0])
		countries := apiResponse[1].([]interface{})
		for i := 2; i <= metadata.Pages; i++ {
			for _, country := range countries {
				result = append(result, *ParseResponse(country))
			}
			client.R().EnableTrace().SetResult(&apiResponse).Get(url + Pagination(i))
			countries = apiResponse[1].([]interface{})
		}

		file, _ := json.MarshalIndent(result, "", " ")
		_ = ioutil.WriteFile(filepath, file, 0644)
	} else {
		file, _ := ioutil.ReadFile(filepath)
		_ = json.Unmarshal([]byte(file), &result)
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
