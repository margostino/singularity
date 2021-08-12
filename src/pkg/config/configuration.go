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
	Enabled       bool
	Population    int
	Countries     int          `yaml:"countries"`
	GenderWeight  GenderWeight `yaml:"gender_weight"`
	CountriesFile string       `yaml:"countries_file"`
	CountriesUrl  string       `yaml:"countries_url"`
}

type Configuration struct {
	PreLoad PreLoad `yaml:"pre_load"`
}

type CommandConfiguration struct {
	Id          string `yaml:"id"`
	Description string `yaml:"description"`
	Args        int    `yaml:"args"`
	Action      string `yaml:"action"`
}

type CommandsConfiguration struct {
	CommandList []CommandConfiguration `yaml:"commands"`
}

var configuration Configuration
var commandsConfiguration CommandsConfiguration

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

func LoadCommandsConfiguration() {
	yamlFile, err := ioutil.ReadFile("../config/commands.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &commandsConfiguration)
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

func GetCountriesFile() string {
	return configuration.PreLoad.CountriesFile
}

func GetCountriesUrl() string {
	return configuration.PreLoad.CountriesUrl
}

func GetPopulation() int {
	return configuration.PreLoad.Population
}

func GetCountries() int {
	return configuration.PreLoad.Countries
}

func GetCommandsConfiguration() *CommandsConfiguration {
	return &commandsConfiguration
}
