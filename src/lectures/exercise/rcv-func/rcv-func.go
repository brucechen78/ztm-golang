//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
	Name      string
}

func (p *Player) AddHealth(health int) {
	p.Health += health
	fmt.Printf("Health added: %v\n", health)
}

func (p *Player) AddEnergy(energy int) {
	p.Energy += energy
	fmt.Printf("Energy added: %v\n", energy)
}

func main() {

	player := Player{
		Health:    100,
		MaxHealth: 100,
		Energy:    100,
		MaxEnergy: 100,
		Name:      "Player",
	}

	fmt.Printf("Player: %+v\n", player)

	player.AddHealth(10)
	player.AddEnergy(10)

	fmt.Printf("Player: %+v\n", player)

}
