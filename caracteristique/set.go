package piscine

import (
	"fmt"
)

func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "sucette" { //On parcours l'inventaire et si on a la sucette on l'utilise
			if p.note > p.notemax-20 { //La sucette rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
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
	marchand := Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "poison": 30, "Skill: go": 60}}
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
		if answer == cle && !p.IsInInventory(answer) {
			if marchand.wallet >= val && p.LimitSpace() {
				marchand.wallet -= val
				delete(marchand.inventaire, cle)
				p.AddInventory(answer)
				fmt.Println("Vous avez ajouté", cle, "à votre inventaire")
				return
			} else {
				fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
				return
			}
		}

	}
	fmt.Printf("%#v n'est pas un objet accessible à la PEDA ou alors tu l'a mal écrit\n", answer)
}
