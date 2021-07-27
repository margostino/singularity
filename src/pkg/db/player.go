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
	Id      string
	Name    string
	Score   uint
	Wallet  uint
	Address Address
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

	var address = NewAddress()
	var newPlayer = NewPlayer(playerName, *address)
	AddNewPlayer(*newPlayer)
}

func NewPlayer(name string, address Address) *Player {
	return &Player{
		Id:      generateId(),
		Name:    name,
		Score:   0,
		Address: address,
		Wallet:  0,
	}
}

func generateId() string {
	return uuid.New().String()
}
