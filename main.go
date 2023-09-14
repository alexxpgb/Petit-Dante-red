package main

import "fmt"

type Personnage struct {
	name       string
	classe     string
	niveau     int
	PDVmax     int
	PDV        int
	inventaire map[string]int
}

func (p *Personnage) Init(name string, classe string) {
	p.name = name
	p.classe = classe
	p.niveau = 1
	p.PDVmax = 100
	p.PDV = 40
	p.inventaire = map[string]int{"potion": 3}
}
func (p Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("Ton nom est :", p.name)
	fmt.Println("Ta classe est :", p.classe)
	fmt.Println("Ton niveau est :", p.niveau)
	fmt.Printf("Tu as %d/%d", p.PDV, p.PDVmax)
	fmt.Println(" point de vie")
	for cle, val := range p.inventaire {
		fmt.Printf("Tu a %d %s", val, cle)
	}
	fmt.Println("-----------------------")
}
func (p Personnage) AccessInventory() {
	fmt.Println("-----------------------")
	fmt.Println("Ton inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("Tu a %d %s", val, cle)
	}
	fmt.Println()
	fmt.Print("-----------------------")
}
func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "potion" {
			if p.PDV > 50 {
				fmt.Println("est tu sur de vouloir utiliser la potion")
			} else {
				p.inventaire["potion"] -= 1
				p.PDV += 50
			}
		} else {
			fmt.Println("Tu n'a pas de potions")
		}
	}
}
func main() {
	var p1 Personnage
	p1.Init("test", "test")
	p1.Display()
	p1.AccessInventory()
}
