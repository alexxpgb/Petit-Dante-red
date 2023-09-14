package piscine

import "fmt"

func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "potion" {
			if p.PDV > p.PDVmax-50 {
				var answer string
				fmt.Println("est tu sur de vouloir utiliser la potion")
				fmt.Scan(&answer)
				if answer == "oui" {
					p.inventaire["potion"] -= 1
					p.PDV = p.PDVmax
					fmt.Println("Tu as pris une potion tu est maintenant à", p.PDV)
				} else if answer == "non" {
					fmt.Println("Fais plus attention la prochaine fois")
				} else {
					fmt.Println("Tu peux répéter ?")
					p.TakePot()
				}
			} else {
				p.inventaire["potion"] -= 1
				p.PDV += 50
				fmt.Println("Tu as pris une potion tu est maintenant à", p.PDV)
			}
		} else {
			fmt.Println("Tu n'a pas de potions")
		}
	}
}

func (p *Personnage) Boutique() {
	marchand := Personnage{name: "Arthur", niveau: 1, PDVmax: 69, PDV: 4, inventaire: map[string]int{"potion": 1}}
	marchand.AccessInventory()
	fmt.Println("Que veux tu parmi tous ses objets")
}
