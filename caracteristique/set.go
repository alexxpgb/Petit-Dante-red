package piscine

import (
	"fmt"
)

func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "sucette" { //On parcours l'inventaire et si on a la sucette on l'utilise
			if p.note > p.notemax-20 { //La sucette rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
				var answer int
				fmt.Println("est tu sur de vouloir utiliser la sucette")
				fmt.Println("1 pour oui")
				fmt.Println("2 pour non")
				fmt.Scan(&answer)
				if answer == 1 {
					p.inventaire["sucette"] -= 1      //Je lui enlève une sucette dans mon inventaire
					if p.inventaire["sucette"] == 0 { //Si j'en ai plus je la supprime de mon inventaire
						delete(p.inventaire, "sucette")
					}
					p.note = p.notemax //Dans le cas où ma note est superieur à ma note max - ma sucette donc ma note est au max car ma sucette fait un +20
					fmt.Println("Tu as pris une sucette tu est maintenant à", p.note)
				} else if answer == 2 {
					fmt.Println("Fais plus attention la prochaine fois")
				} else {
					fmt.Println("Tu peux répéter ?")
					p.TakePot() //On le relance
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
	fmt.Println("Tu n'as pas de sucette") // j'ai regardé partout dans mon inventaire mais tu n'as pas de sucette
}

func (p *Personnage) Boutique() {
	marchand := Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "poison": 30, "Skill: go": 60}} //L'inventaire d'Arthur est la boutique
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("%s pour %d€", cle, val) //La clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Println("Que veux tu parmi tous ses objets")
	var answer string
	fmt.Scan(&answer)
	for cle, val := range marchand.inventaire {
		if answer == cle { //Je parcoure mon inventaire et je vérifie que l'objet que l'utilisateur ma selectione est bien dans mon inventaire
			if marchand.wallet >= val && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
				marchand.wallet -= val //Je lui prends l'argent
				if cle == "sucette" {  //Car la première sucette est gratuite et après c'est payant
					delete(marchand.inventaire, cle)
					marchand.inventaire["sucette"] = 20
				}
				p.AddInventory(answer) //Je l'ajoute dans mon inventaire
				fmt.Println("Vous avez ajouté", cle, "à votre inventaire")
				return
			} else {
				fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
				return
			}
		}
		if answer[:4] == "Skill" && answer == cle { //Pour le cas ou il achète un skill
			if marchand.wallet >= val && p.LimitSpace() {
				marchand.wallet -= val
				delete(marchand.inventaire, cle) //La je dois le supprimer de l'inventaire de mon marchand
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
	p.Menu()
}
