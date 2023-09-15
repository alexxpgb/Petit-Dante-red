package piscine

import "fmt"

func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "sucette" {
			if p.note > p.notemax-20 {
				var answer string
				fmt.Println("est tu sur de vouloir utiliser la sucette")
				fmt.Scan(&answer)
				if answer == "oui" {
					p.inventaire["sucette"] -= 1
					if p.inventaire["sucette"] == 0 {
						delete(p.inventaire, "sucette")
					}
					p.note = p.notemax
					fmt.Println("Tu as pris une sucette tu est maintenant à", p.note)
				} else if answer == "non" {
					fmt.Println("Fais plus attention la prochaine fois")
				} else {
					fmt.Println("Tu peux répéter ?")
					p.TakePot()
				}
			} else {
				p.inventaire["sucette"] -= 1
				p.note += 20
				fmt.Println("Tu as pris une sucette tu est maintenant à", p.note)
			}
		} else {
			fmt.Println("Tu n'a pas de sucette")
		}
	}
}

func (p *Personnage) Boutique() {
	marchand := Personnage{name: "Arthur", niveau: "B1", notemax: 100, note: 4, inventaire: map[string]int{"multiprise": 1}}
	marchand.AccessInventory()
	fmt.Println("Que veux tu parmi tous ses objets")
	var answer string
	fmt.Scan(&answer)
	for cle := range marchand.inventaire {
		if answer == cle {
			marchand.inventaire[cle] -= 1
			p.inventaire[cle] = 1
		}
	}
}
