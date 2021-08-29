package db

import "math/rand"

var Players []Player
var Countries []Country

func AddNewPlayer(player Player) {
	Players = append(Players, player)
}

func GetPlayerBy(username string) *Player {
	for _, v := range Players {
		if v.Username == username {
			return &v
		}
	}
	return nil
}

func GetCountryBy(name string) Country {
	var country Country
	for k, v := range Countries {
		if Countries[k].Name == name {
			country = v
			break
		}
	}
	return country
}

func AddNewCountry(country Country) {
	Countries = append(Countries, country)
}

func PickAddress() *Address {
	indexCountry := rand.Intn(len(Countries) - 1)
	country := Countries[indexCountry]
	indexAddress := rand.Intn(len(country.Addresses) - 1)
	address := country.Addresses[indexAddress]
	return &address
}

func PickCountry() *Country {
	indexCountry := rand.Intn(len(Countries) - 1)
	return &Countries[indexCountry]
}

func PickPlayer() *Player {
	indexPlayer := rand.Intn(len(Players) - 1)
	return &Players[indexPlayer]
}

func GetCountries() *[]Country {
	return &Countries
}
