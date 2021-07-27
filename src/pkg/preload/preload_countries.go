package preload

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"org.gene/singularity/pkg/db"
	"strconv"
	"time"
)

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

func getCountries() []db.Country {
	var result []db.Country
	type CountryProviderResponse interface{}
	apiResponse := make([]CountryProviderResponse, 0)
	client := resty.New()
	//response, err := client.R().EnableTrace().SetResult(&apiResponse).Get("https://api.worldbank.org/v2/country?format=json")
	client.R().EnableTrace().SetResult(&apiResponse).Get("https://api.worldbank.org/v2/country?format=json")
	//metadata := apiResponse[0].(map[string]interface{})
	countries := apiResponse[1].([]interface{})

	for _, value := range countries {
		countryMap := value.(map[string]interface{})
		regionMap := countryMap["region"].(map[string]interface{})
		region := db.Region{Id: regionMap["id"].(string), Name: regionMap["value"].(string)}
		latitude, _ := strconv.ParseFloat(countryMap["latitude"].(string), 64)
		longitude, _ := strconv.ParseFloat(countryMap["longitude"].(string), 64)
		addresses := getAddresses()
		country := db.Country{Id: countryMap["id"].(string),
			Name:      countryMap["name"].(string),
			Capital:   countryMap["capitalCity"].(string),
			Latitude:  latitude,
			Longitude: longitude,
			Region:    region,
			Addresses: addresses,
		}
		result = append(result, country)
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
