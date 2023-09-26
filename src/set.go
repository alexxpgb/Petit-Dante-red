package piscine

import (
	"fmt"
	"strconv"
)

// Je l'initialise ici pour que ce soit une variable globale et qu'elle puisse etre modifiée(pour sucette)
var marchand Personnage = Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "douche": 15, "Skill: go": 20, "Fourrure de loup": 5, "Plume de Corbeau": 5, "Cuir de Sanglier": 5, "Peau de Troll": 5, "Upgrade inventaire": 20, "skittles": 15}}
var forgeron Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"casque gaming": 5, "pull ynov": 5, "multiprice": 5}}

func (p *Personnage) Boutique() {
	lst := TransvalseList(marchand.inventaire)
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, marchand.inventaire[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Print("❖ Que veux tu parmi tous ses objets\n\n\n\n\n\n")
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
			if string(lst[i])[:4] == "Skill" {
				delete(marchand.inventaire, string(lst[i]))
			}
			p.AddInventory(string(lst[i])) //Je l'ajoute dans mon inventaire
			fmt.Println("Vous avez ajouté", string(lst[i]), "à votre inventaire")
			p.Menu()
		} else {
			fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
			p.Menu() //Je reviens au menu dans tous les cas
		}
	}
	fmt.Printf("%#v J'attendais le numero de l'objet que tu voulais acheter\n", answer)
	p.Menu()
}

func (p *Personnage) TakeInt(nb int) { //TakePot pour le mana
	var a bool
	for cle := range p.inventaire {
		if cle == "skittles" { //On parcours l'inventaire et si on a la skittles on l'utilise
			if p.energy > int(p.intmax)-20 { //La skittles rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
				fmt.Println("❖ Est tu sur de vouloir utiliser un skittles")
				fmt.Println("1 pour oui")
				fmt.Println("2 pour non")
				switch Scan() {
				case "1":
					p.inventaire["skittles"] -= 1      //Je lui enlève une skittles dans mon inventaire
					if p.inventaire["skittles"] == 0 { //Si j'en ai plus je la supprime de mon inventaire
						delete(p.inventaire, "skittles")
					}
					p.energy = int(p.intmax) //Dans le cas où ma note est superieur à ma note max - ma skittles donc ma note est au max car ma skittles fait un +20
					fmt.Println("Tu as pris une skittles tu à maintenant", p.energy, "energie")
					a = true
				case "2":
					fmt.Println("Fais plus attention la prochaine fois")
					a = true
				default:
					fmt.Println("Tu peux répéter ?")
					p.TakeInt(nb) //On le relance

				}
			} else {
				p.inventaire["skittles"] -= 1
				if p.inventaire["skittles"] == 0 {
					delete(p.inventaire, "skittles")
				}
				p.energy += 20
				fmt.Println("Tu as pris une skittles tu à maintenant ", p.energy, "energie")
				a = true
			}
		}
	}
	if !a { //Le cas contraire de a donc si on utilise une skittles on peut pas rentrer dans cette boucle
		fmt.Println("Tu n'as pas de skittles") // j'ai regardé partout dans mon inventaire mais tu n'as pas de skittles
	}
	if nb == 2 {
		return
	}
}

func (p *Personnage) Forgeron() {
	lst := TransvalseList(forgeron.inventaire)
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire du forgeron est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, forgeron.inventaire[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Print("❖ Que veux tu parmi tous ses objets\n\n\n\n")
	answer := Scan() //Le mec il rentre le numero auquel est attribué à son objet
	i, _ := strconv.Atoi(answer)
	i -= 1
	if p.wallet >= forgeron.inventaire[lst[i]] && p.LimitSpace() {
		switch answer {
		case "1":
			if p.IsInInventory("Plume de Corbeau") && p.IsInInventory("Cuir de Sanglier") {
				p.RemoveInventory("Plume de Corbeau")
				p.RemoveInventory("Cuir de Sanglier")
				p.AddInventory("casque gaming")
				p.wallet -= forgeron.inventaire[lst[i]]
				p.Menu()
			}
		case "2":
			if p.IsInInventory("Fourrure de loup") && p.IsInInventory("Peau de Troll") {
				p.RemoveInventory("Peau de Troll")
				p.RemoveInventory("Fourrure de loup")
				if p.IsInInventory(" Fourrure de loup") {
					p.RemoveInventory("Fourrure de loup")
					p.AddInventory("pull ynov")
					p.wallet -= forgeron.inventaire[lst[i]]
					p.Menu()
				} else {
					fmt.Println("tu n'as pas de quoi fabriquer cet objet")
					p.AddInventory("Fourrure de loup")
					p.AddInventory("Peau de Troll")
					p.Menu()
				}
			}
		case "3":
			if p.IsInInventory("Fourrure de loup") && p.IsInInventory("Cuir de Sanglier") {
				p.RemoveInventory("Fourrure de loup")
				p.RemoveInventory("Cuir de Sanglier")
				p.AddInventory("multiprice")
				p.wallet -= forgeron.inventaire[lst[i]]
				p.Menu()
			}
		}
	}
}
