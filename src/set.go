package piscine

import (
	"fmt"
	"strconv"
)

// Je l'initialise ici pour que ce soit une variable globale et qu'elle puisse etre modifiée(pour sucette)
<<<<<<< HEAD
var marchand Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"sucette": 0, "douche": 15, "Skill: go": 20, "casque gaming": 5, "pull ynov": 5, "multiprice": 5, "Upgrade inventaire": 20, "skittles": 15}}
=======
var marchand Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"sucette": 0, "poison": 15, "Skill: go": 20, "Chapeau de l’aventurier": 5, "Tunique de l’aventurier": 5, "Bottes de l’aventurier": 5, "Upgrade inventaire": 20}}
>>>>>>> 4590ea0a91dec6bfa77de47ea9f0c0c968b52d69

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
	lst := TransvalseList(marchand.inventaire)
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, marchand.inventaire[val]) //le count il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
<<<<<<< HEAD
	fmt.Print("❖ Que veux tu parmi tous ses objets\n\n\n\n\n\n")
	answer := Scan()                                      //Le mec il rentre le numero auquel est attribué son objet
	i, _ := strconv.Atoi(answer)                          //string en int
	i -= 1                                                //Pour avoir un indice de 0
	if answer <= strconv.Itoa(len(lst)) && answer > "0" { //On vérifie que l'utilisateur a bien rentré qlq chose compris entre 1 et la taille de ma liste
		if p.wallet >= marchand.inventaire[lst[i]] && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
			p.wallet -= marchand.inventaire[lst[i]] //Je lui prends l'argent
			if string(lst[i]) == "sucette" {        //Car la première sucette est gratuite et après c'est payant
=======
	fmt.Println("❖ Que veux tu parmi tous ses objets")
	answer := Scan() //Le mec il rentre le numero auquel est attribué à son objet
	i, _ := strconv.Atoi(answer)
	i -= 1
	if answer < strconv.Itoa(len(lst)) && answer > "0" { //On vérifie que l'utilisateur a bien rentré
		if p.wallet >= p.inventaire[lst[i]] && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
			p.wallet -= p.inventaire[lst[i]] //Je lui prends l'argent
			if string(lst[i]) == "sucette" { //Car la première sucette est gratuite et après c'est payant
>>>>>>> 4590ea0a91dec6bfa77de47ea9f0c0c968b52d69
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
	lst := TransvalseList(marchand.inventaire)
	forgeron := Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"Chapeau de l’aventurier": 5, "Tunique de l’aventurier": 5, "Bottes de l’aventurier": 5}}
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire du forgeron est composé de")
	for cle, val := range forgeron.inventaire {
		fmt.Printf("๑%s pour %d € ;\n", cle, val)
	} 
	fmt.Println("\n----------------------")
	fmt.Println("❖ Que veux tu parmi tous ses objets")
	answer := Scan() //Le mec il rentre le numero auquel est attribué à son objet
	i, _ := strconv.Atoi(answer)
	i -= 1
	if p.wallet >= p.inventaire[lst[i]] && p.LimitSpace(){
			switch asnwer {
	case "Chapeau de l’aventurier" :
		if p.IsInInventory("Plume de Corbeau") && p.IsInInventory("Cuir de Sanglier"){
			p.RemoveInventory("Plume de Corbeau")
			p.RemoveInventory("Cuir de Sanglier")
			p.AddInventory("Chapeau de l’aventurier")
			p.wallet -= 5
			p.Menu()
		}
	case "Tunique de l’aventurier":
		if p.IsInInventory(" Fourrure de loup") && p.IsInInventory("Peau de Troll"){
			p.RemoveInventory("Peau de Troll")
			p.RemoveInventory("Fourrure de loup")
			if p.IsInInventory(" Fourrure de loup"){
				p.RemoveInventory("Fourrure de loup")
				p.AddInventory("Tunique de l’aventurier")
				p.wallet -= 5
				p.Menu()
			}else{
				fmt.Println("tu n'as pas de quoi fabriquer cet objet")
				p.AddInventory("Fourrure de loup")
				p.AddInventory("Peau de Troll")
				p.wallet -= 5
				p.Menu()
			}
		case "Bottes de l’aventurier":
			if p.IsInInventory("Fourrure de loup") && p.IsInInventory("Cuir de Sanglier"){
				p.RemoveInventory("Fourrure de loup")
				p.RemoveInventory("Cuir de Sanglier")
				p.AddInventory("Bottes de l’aventurier")
				p.wallet -= 5
				p.Menu()
			}
	}

	
			
}
  
