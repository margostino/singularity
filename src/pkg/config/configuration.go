package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type GenderWeight struct {
	Male, Female uint
}

type PreLoad struct {
	Enabled           bool
	Population        int
	CountriesQuantity int
	GenderWeight      GenderWeight `yaml:"gender_weight"`
}

type Configuration struct {
	PreLoad PreLoad `yaml:"pre_load"`
}

var configuration Configuration

func LoadConfiguration() {
	yamlFile, err := ioutil.ReadFile("../config/configuration.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func IsPreLoadEnabled() bool {
	return configuration.PreLoad.Enabled
}

func GetGenderWeight() (uint, uint) {
	return configuration.PreLoad.GenderWeight.Male, configuration.PreLoad.GenderWeight.Female
}

func GetPopulation() int {
	return configuration.PreLoad.Population
}

func GetCountries() int {
	return configuration.PreLoad.CountriesQuantity
}
