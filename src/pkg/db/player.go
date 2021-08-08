package db

import (
	"bufio"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"os"
	"strings"
)

type Player struct {
	Id       string
	Name     string
	Username string
	Score    uint
	Wallet   uint
	Address  Address
}

type Gender struct {
	name  string
	value int
}

var Male = Gender{
	name: "male", value: randomdata.Male,
}

var Female = Gender{
	name: "female", value: randomdata.Female,
}

func CreatePlayer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name?")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.Replace(playerName, "\n", "", -1)
	var username = GetUsername(playerName)
	var address = NewAddress()
	var newPlayer = NewPlayer(playerName, username, *address)
	AddNewPlayer(*newPlayer)
}

func GetUsername(name string) string {
	lowercase := strings.ToLower(name)
	names := strings.Split(lowercase, " ")
	return strings.Join(names, "_")
}

func NewPlayer(name string, username string, address Address) *Player {
	return &Player{
		Id:       generateId(),
		Name:     name,
		Username: username,
		Score:    0,
		Address:  address,
		Wallet:   0,
	}
}

func generateId() string {
	return uuid.New().String()
}
