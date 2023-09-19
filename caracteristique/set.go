package piscine

import (
	"fmt"
)

// Je l'initialise ici pour que ce soit une variable globale et qu'elle puisse etre modifiée(pour sucette)
var marchand Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"sucette": 0, "poison": 15, "Skill: go": 20, "Chapeau de l’aventurier": 5, "Tunique de l’aventurier": 5, "Bottes de l’aventurier": 5, "Upgrade inventaire": 20}}

func (p *Personnage) TakePot() {
	for cle := range p.inventaire {
		if cle == "sucette" { //On parcours l'inventaire et si on a la sucette on l'utilise
			if p.note > p.notemax-20 { //La sucette rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
				var answer int
				fmt.Println("❖ Est tu sur de vouloir utiliser la sucette")
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
	var count int = 1
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for cle, val := range marchand.inventaire {
		fmt.Printf("๑ %d %s pour %d € \n", count, cle, val) //le count il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
		count++
	}
	fmt.Println("\n----------------------")
	fmt.Println("❖ Que veux tu parmi tous ses objets")
	var answer int
	fmt.Scan(&answer)
	count = 1 //Le mec il rentre le numero auquel est attribuéson objet
	for cle, val := range marchand.inventaire {
		if answer == count && count != 3 { //Je selectionne l'objet en question sauf si la 3 car c'est un skill et il se passe pas la meme chose
			if marchand.wallet >= val && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
				marchand.wallet -= val //Je lui prends l'argent
				if cle == "sucette" {  //Car la première sucette est gratuite et après c'est payant
					delete(marchand.inventaire, cle)
					marchand.inventaire["sucette"] = 20
				}
				p.AddInventory(cle) //Je l'ajoute dans mon inventaire
				fmt.Println("Vous avez ajouté", cle, "à votre inventaire")
				p.Menu()
			} else {
				fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
				p.Menu() //Je reviens au menu dans tous les cas
			}
		}
		if answer == 3 { //Pour le cas ou il achète un skill
			if marchand.wallet >= val && p.LimitSpace() {
				marchand.wallet -= val
				delete(marchand.inventaire, cle) //La je dois le supprimer de l'inventaire de mon marchand
				p.AddInventory(cle)
				fmt.Println("Vous avez ajouté", cle, "à votre inventaire")
				p.Menu()
			} else {
				fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
				p.Menu()
			}
		}
		count++
	}
	fmt.Printf("%#v J'attendais le numero de l'objet que tu voulais acheter\n", answer)
	p.Menu()
}

func (p *Personnage) Forgeron() {
	forgeron := Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"Chapeau de l’aventurier": 5, "Tunique de l’aventurier": 5, "Bottes de l’aventurier": 5}}
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire du forgeron est composé de")
	for cle, val := range forgeron.inventaire {
		fmt.Printf("๑%s pour %d € ;\n", cle, val)
	}
}
