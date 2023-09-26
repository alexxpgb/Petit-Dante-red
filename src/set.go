package piscine

import (
	"fmt"
	"strconv"
)

// Je l'initialise ici pour que ce soit une variable globale et qu'elle puisse etre modifiée(pour sucette)
var AllOfObject map[string]int = map[string]int{"sucette": 20, "douche": 15, "Skill: go": 20, "febreze": 5, "carte graphique": 5, "souris": 5, "clavier mecanique": 5, "Upgrade inventaire": 20, "skittles": 15, "pull ynov": 25, "multiprice": 25, "casque gaming": 25}
var marchand Personnage = Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "douche": 15, "Skill: go": 20, "febreze": 5, "carte graphique": 5, "souris": 5, "clavier mecanique": 5, "Upgrade inventaire": 20, "skittles": 15}}
var forgeron Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"casque gaming": 15, "pull ynov": 15, "multiprice": 15}}

func (p *Personnage) Boutique() {
	lst := TransvalseList(marchand.inventaire)
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, marchand.inventaire[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Print("❖ Veux tu quelque chose parmi tous ses objets ou peut être vendre\n1-Acheter\n2-Vendre\n\n\n\n\n")
	answer := Scan()
	if answer == "1" {
		fmt.Print("Quel objet veut tu ?\n\n\n\n")
		answer = Scan()                                       //Le mec il rentre le numero auquel est attribué son objet
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
	} else if answer == "2" {
		lst = TransvalseList(p.inventaire)
		fmt.Println("-----------------------")
		fmt.Println("Que veut tu vendre parmi tous ses objets")
		for ind, val := range lst {
			fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, AllOfObject[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
		}
		fmt.Print("\n----------------------\n\n\n\n\n")
		answer = Scan()              //Le mec il rentre le numero auquel est attribué son objet
		i, _ := strconv.Atoi(answer) //string en int
		i -= 1
		if answer <= strconv.Itoa(len(lst)) && answer > "0" {
			if IsInMap(AllOfObject, string(lst[i])) {
				p.wallet += AllOfObject[(lst[i])]
				marchand.inventaire[lst[i]+" reconditionée"] = AllOfObject[lst[i]] * 2
			} else {
				fmt.Println("Le marchand ne peut pas t'acheter cette objet")
				p.Menu()

			}
			p.inventaire[lst[i]]--
			if p.inventaire[lst[i]] == 0 {
				delete(p.inventaire, string(lst[i]))
			}
			fmt.Printf("Vous avez bien vendu votre objet pour %d€\n", AllOfObject[lst[i]])
			p.Menu()
		}
	}
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
			if p.IsInInventory("carte graphique") && p.IsInInventory("souris") {
				p.RemoveInventory("carte graphique")
				p.RemoveInventory("souris")
				p.AddInventory("casque gaming")
				p.wallet -= forgeron.inventaire[lst[i]]
				p.Menu()
			}
		case "2":
			if p.IsInInventory("febreze") && p.IsInInventory("clavier mecanique") {
				p.RemoveInventory("clavier mecanique")
				p.RemoveInventory("febreze")
				if !p.IsInInventory(" febreze") {
					fmt.Println("tu n'as pas de quoi fabriquer cet objet")
					p.AddInventory("febreze")
					p.AddInventory("clavier mecanique")
					fmt.Println("tu possèdes maintenant un clavier mecanique")
					p.Menu()
				} else {

					p.RemoveInventory("febreze")
					p.AddInventory("pull ynov")
					fmt.Println("tu possèdes maintenant un pull ynov ")
					p.wallet -= forgeron.inventaire[lst[i]]
					p.Menu()
				}
			}
		case "3":
			if p.IsInInventory("febreze") && p.IsInInventory("souris") {
				p.RemoveInventory("febreze")
				p.RemoveInventory("souris")
				p.AddInventory("multiprise")
				fmt.Println("tu possèdes maintenant une multiprise")
				p.wallet -= forgeron.inventaire[lst[i]]
				p.Menu()
			}
		default:
			fmt.Println("je n'ai pas compris ta répone, peux tu repeter ? ")
			p.Forgeron()
		}
	} else {
		fmt.Println("tu ne possèdes pas suffisaement d'argent ou ton invnetiare est plein. reviens plus tard . ")
		p.Forgeron()
	}
}

func (p *Personnage) Garden() {
	fmt.Println("Bienvenue dans la terrasse")
}
