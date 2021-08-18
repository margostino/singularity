package db

import (
	"fmt"
)

type Region struct {
	Id   string
	Name string
}

type Country struct {
	Id             string
	Name           string
	Capital        string
	Latitude       float64
	Longitude      float64
	Region         Region
	WarmingMetrics []WarmingMetric
	Addresses      []Address
}

type WarmingMetric struct {
	Key         string
	Value       float64
	Unit        string
	Description string
}

type Address struct {
	Street    string
	City      string
	State     string
	Zip       string
	Country   string
	Latitude  float64
	Longitude float64
}

func NewAddress() *Address {
	var latitude, longitude float64
	var street, city, state, zip, country string

	fmt.Println("Street?")
	fmt.Scanf("%d", &street)
	fmt.Println("City?")
	fmt.Scanf("%d", &city)
	fmt.Println("State?")
	fmt.Scanf("%d", &state)
	fmt.Println("Zip?")
	fmt.Scanf("%d", &zip)
	fmt.Println("Country?")
	fmt.Scanf("%d", &country)
	fmt.Println("Address Latitude?")
	fmt.Scanf("%d", &latitude)
	fmt.Println("Address Longitude?")
	fmt.Scanf("%d", &longitude)

	return &Address{
		street,
		city,
		state,
		zip,
		country,
		latitude,
		longitude,
	}
}
