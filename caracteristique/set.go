package piscine

import (
	"fmt"
)

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
				return
			} else {
				p.inventaire["sucette"] -= 1
				if p.inventaire["sucette"] == 0 {
					delete(p.inventaire, "sucette")
				}
				p.note += 20
				fmt.Println("Tu as pris une sucette tu est maintenant à", p.note)
				return
			}
		}
	}
	fmt.Println("Tu n'as pas de sucette")
}

func (p *Personnage) Boutique() {
	marchand := Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "poison": 30}}
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("%s pour %d€", cle, val)
	}
	fmt.Println("\n----------------------")
	fmt.Println("Que veux tu parmi tous ses objets")
	var answer string
	fmt.Scan(&answer)
	for cle, val := range marchand.inventaire {
		if answer == cle {
			if marchand.wallet >= val {
				marchand.wallet -= val
				delete(marchand.inventaire, cle)
				p.AddInventory(answer)
				return
			} else {
				fmt.Println("Pas sur que tu peux te payer ça")
			}
		}
	}
	fmt.Printf("%#v n'est pas un objet accessible à la PEDA\n", answer)
}
