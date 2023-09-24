package piscine

import (
	"fmt"
	"strconv"
)

// Je l'initialise ici pour que ce soit une variable globale et qu'elle puisse etre modifiée(pour sucette)
var marchand Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"sucette": 0, "douche": 15, "Skill: go": 20, "casque gaming": 5, "pull ynov": 5, "multiprice": 5, "Upgrade inventaire": 20}}

// casque gaming =chapeau de l'aventurier; pull ynov =tunique de l'aventurier ; multiprice=Bottes de l’aventurier

//J'ai mit le take pot dans le dead plus logique

func (p *Personnage) Boutique() {
	lst := TransvalseList(marchand.inventaire) //Je cast ma map en liste
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, marchand.inventaire[val]) //le ind il sert juste pour numeroter les items la val est le nom de l'objet et marchand.inventaire[val] est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Println("❖ Que veux tu parmi tous ses objets\n\n\n\n\n")
	answer := Scan()                                      //Le mec il rentre le numero auquel est attribué son objet
	i, _ := strconv.Atoi(answer)                          //string en int
	i -= 1                                                //Pour avoir un indice de 0
	if answer <= strconv.Itoa(len(lst)) && answer > "0" { //On vérifie que l'utilisateur a bien rentré qlq chose compris entre 1 et la taille de ma liste
		if p.wallet >= marchand.inventaire[lst[i]] && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
			p.wallet -= marchand.inventaire[lst[i]] //Je lui prends l'argent
			if string(lst[i]) == "sucette" {        //Car la première sucette est gratuite et après c'est payant
				delete(marchand.inventaire, string(lst[i]))
				marchand.inventaire["sucette"] = 20
			}
			if string(lst[i])[:4] == "Skill" { //On supprime quand c'est un skill
				delete(marchand.inventaire, string(lst[i]))
			}
			p.AddInventory(string(lst[i])) //Je l'ajoute dans mon inventaire
			fmt.Println("Vous avez ajouté", string(lst[i]), "à votre inventaire")
			p.Menu() //tjr rappelé le menu() sinon fin du programme
		} else {
			fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
			p.Menu() //Je reviens au menu dans tous les cas
		}
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
