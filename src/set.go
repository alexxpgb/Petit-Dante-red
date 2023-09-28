package piscine

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

// J'initialise toute mes variables pour qu'elles soit globale et qu'elle puisse etre modifiée
var AllOfObject map[string]int = map[string]int{"sucette": 20, "douche": 15, "Skill: go": 20, "febreze": 5, "carte graphique": 5, "souris": 5, "clavier mecanique": 5, "Upgrade inventaire": 20, "skittles": 15, "pull ynov": 20, "multiprise": 15, "casque gaming": 15}
var marchand Personnage = Personnage{name: "Arthur", classe: "marchand", inventaire: map[string]int{"sucette": 0, "douche": 15, "Skill: go": 20, "febreze": 5, "carte graphique": 5, "souris": 5, "clavier mecanique": 5, "Upgrade inventaire": 20, "skittles": 15}}
var forgeron Personnage = Personnage{name: "Vito", classe: "forgeron", inventaire: map[string]int{"casque gaming": 15, "pull ynov": 20, "multiprise": 15}}
var jardin map[string]int = map[string]int{}
var min int

func (p *Personnage) Boutique() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\n\n\n\n\n\n\n\n")
	TermPrint("______        _                                                ", 1, 1, termbox.ColorCyan)
	TermPrint("| ___ \\      | |                                              ", 1, 2, termbox.ColorCyan)
	TermPrint("| |_/ /__  __| | __ _                                          ", 1, 3, termbox.ColorCyan)
	TermPrint("|  __/ _ \\/ _` |/ _` |                                        ", 1, 4, termbox.ColorCyan)
	TermPrint("| | |  __/ (_| | (_| |                                         ", 1, 5, termbox.ColorCyan)
	TermPrint("\\_|  \\___|\\__,_|\\__,_|                                     ", 1, 6, termbox.ColorCyan)

	lst := TransvalseList(marchand.inventaire) //Une map se modifie sans arrêt une liste c'est plus pratique
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire de la PEDA est composé de")
	for ind, val := range lst {
		fmt.Printf("• %d/ %s pour %d € \n", ind+1, val, marchand.inventaire[val]) //le ind il sert juste pour numeroter les items la val est le nom de l'objet et marchand.inventaire[val] est le montant
	}
	fmt.Print("\n----------------------\n\n\n")
	fmt.Print("❖ Veux tu quelque chose parmi tous ses objets ou peut être vendre\n1-Acheter\n2-Vendre\n\n\n\n")
	answer := Scan()
	if answer == "1" {
		fmt.Print("Quel objet veut tu ?\n\n\n\n\n")
		ans := Scan()               //Le mec il rentre le numero auquel est attribué son objet
		i, err := strconv.Atoi(ans) //string en int
		i -= 1                      //Pour avoir un indice de 0
		if err != nil {             //Il m'a pas mit un numero
			fmt.Println("Pas compris")
			Enter()
			p.Boutique()
		}
		if ans <= strconv.Itoa(len(lst)) && ans > "0" { //On vérifie que l'utilisateur a bien rentré qlq chose compris entre 1 et la taille de ma liste
			if p.wallet >= marchand.inventaire[lst[i]] && p.LimitSpace() { //Je vérifie si j'ai assez d'argent dans mon portefeuille et que j'ai la place dans mon inventaire
				p.wallet -= marchand.inventaire[lst[i]] //Je lui prends l'argent
				if string(lst[i]) == "sucette" {        //Car la première sucette est gratuite et après c'est payant
					delete(marchand.inventaire, string(lst[i]))
					marchand.inventaire["sucette"] = 20
				}
				if string(lst[i])[:5] == "Skill" { //Avec les skills c'est pas la même
					delete(marchand.inventaire, string(lst[i]))
				}
				p.AddInventory(string(lst[i])) //Je l'ajoute dans mon inventaire
				fmt.Println("Vous avez ajouté", string(lst[i]), "à votre inventaire")
				Enter()
				p.Menu()
			} else {
				fmt.Println("Pas sur que tu peux te payer ça ou tu n'as pas assez de place dans ton inventaire")
				Enter()
				p.Menu() //Je reviens au menu dans tous les cas
			}
		}
		fmt.Printf("%v J'attendais le numero de l'objet que tu voulais acheter\n", ans)
		Enter()
		p.Menu()
	} else if answer == "2" {
		if !p.IsInInventory("casque gaming") && !p.IsInInventory("pull ynov") && !p.IsInInventory("multiprise") {
			fmt.Println("La peda n'accepte que les équipement ")
			Enter()
			p.Menu()
		}
		lst = TransvalseList(p.inventaire)
		fmt.Println("-----------------------")
		fmt.Println("Veux tu vendre quelque chose ou recycler\n1/vendre \n2/Recycler\n")
		var count int = 1
		lst2 := []string{}
		for _, val := range lst {
			if val == "casque gaming" || val == "pull ynov" || val == "multiprise" {
				fmt.Printf("๑ %d/ %s pour %d € \n", count, val, AllOfObject[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
				count++
				lst2 = append(lst2, val)
			}
		}
		fmt.Print("\n----------------------\n\n\n\n\n")
		rep := Scan()
		switch rep {
		case "1":
			fmt.Print("Quel objet veux tu vendre?\n\n\n\n")
			ans := Scan()
			i, err := strconv.Atoi(ans)
			i -= 1
			if err != nil {
				fmt.Println("Pas compris")
				Enter()
				p.Boutique()
			}

			if ans <= strconv.Itoa(len(lst2)) && ans > "0" {
				if IsInMap(AllOfObject, string(lst2[i])) { //AllOfObject contient tous les items avec leurs prix
					p.wallet += AllOfObject[(lst2[i])]
				} else {
					fmt.Println("Le marchand ne peut pas t'acheter cette objet")
					Enter()
					p.Menu()

				}
				p.inventaire[lst2[i]]--
				if p.inventaire[lst2[i]] == 0 {
					delete(p.inventaire, string(lst2[i]))
				}
				fmt.Printf("Vous avez bien vendu votre objet pour %d€\n\n\n", AllOfObject[lst2[i]])
				Enter()
				p.Menu()
			} else {
				fmt.Println("Pas compris")
				Enter()
				p.Boutique()
			}

		case "2":
			lst = TransvalseList(p.inventaire)
			fmt.Println("-----------------------")
			fmt.Println("Parmi ses objets lequel veux tu recycler")
			var count int = 1
			lst = []string{}
			for _, val := range lst {
				if val == "casque gaming" || val == "pull ynov" || val == "multiprise" {
					fmt.Printf("๑ %d/ %s pour %d € \n", count, val, AllOfObject[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
					count++
					lst2 = append(lst2, val)
				}
			}
			fmt.Print("\n----------------------\n\n\n\n\n")
			ans := Scan()
			i, err := strconv.Atoi(ans)
			i -= 1
			if err != nil {
				fmt.Println("Pas compris")
				Enter()
				p.Boutique()
			}
			fmt.Println(ans, strconv.Itoa(len(lst2)))
			if ans <= strconv.Itoa(len(lst2)) && ans > "0" && p.LimitSpace() {
				if lst2[i] == "multiprise" {
					if rand.Float64() < 5.0 {
						p.AddInventory("febreze")
						fmt.Println("Bravo tu as gagné un febreze")
						Enter()
						p.Menu()
					} else {
						p.AddInventory("souris")
						fmt.Println("Bravo tu as gagné une souris")
						Enter()
						p.Menu()
					}
				} else if lst2[i] == "pull ynov" {
					if rand.Float64() < 5.0 {
						p.AddInventory("febreze")
						fmt.Println("Bravo tu as gagné un febreze")
						Enter()
						p.Menu()
					} else {
						p.AddInventory("clavier mecanique")
						fmt.Println("Bravo tu as gagné un clavier mecanique")
						Enter()
						p.Menu()
					}
				} else if lst2[i] == "casque gaming" {
					if rand.Float64() < 5.0 {
						p.AddInventory("souris")
						fmt.Println("Bravo tu as gagné une souris")
						Enter()
						p.Menu()
					} else {
						p.AddInventory("carte graphique")
						fmt.Println("Bravo tu as gagné une carte graphique")
						Enter()
						p.Menu()
					}
				}
			} else {
				fmt.Println("Vous n'avez pas selectionner l'index indiqué ou vous n'avez pas la place nécessaire")
			}
		}
	}
	fmt.Println("Pas compris")
	Enter()
	p.Boutique()
}

func (p *Personnage) TakeInt(nb int) { //TakePot pour le mana
	var a bool
	for cle := range p.inventaire {
		if cle == "skittles" { //On parcours l'inventaire et si on a la skittles on l'utilise
			if p.energy > int(p.intmax)-20 { //La skittles rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
				fmt.Println("❖ Est tu sur de vouloir utiliser un skittles")
				fmt.Println("1 pour oui")
				fmt.Print("2 pour non\n\n\n\n")
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
	fmt.Print("\033[H\033[2J")
	TermPrint("    ___       __          _                                              	", 1, 1, termbox.ColorCyan)
	TermPrint("   /   | ____/ /___ ___  (_)___                                            	", 1, 2, termbox.ColorCyan)
	TermPrint("  / /| |/ __  / __ `__ \\/ / __ \\                                       	", 1, 3, termbox.ColorCyan)
	TermPrint(" / ___ / /_/ / / / / / / / / / /                                            	", 1, 4, termbox.ColorCyan)
	TermPrint("/_/  |_\\__,_/_/ /_/ /_/_/_/ /_/ 	                                        ", 1, 5, termbox.ColorCyan)
	fmt.Print("\n\n\n\n\n\n")
	lst := TransvalseList(forgeron.inventaire)
	fmt.Println("-------------------------------------")
	fmt.Println("L'inventaire du forgeron est composé de (partir=0)")
	TermPrint("• 1 casque gaming pour 15 € ", 3, 8, termbox.ColorDefault) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	TermPrint("• 2 pull ynov pour 15 € ", 3, 9, termbox.ColorDefault)
	TermPrint("• 3 multiprise pour 15 € ", 3, 10, termbox.ColorDefault)

	if p.IsInInventory("carte graphique") {
		TermPrint("carte graphique", 30, 8, termbox.ColorGreen)
		if p.IsInInventory("souris") {
			TermPrint("souris", 48, 8, termbox.ColorGreen)
		} else {
			TermPrint("souris", 48, 8, termbox.ColorRed)
		}
	} else {
		TermPrint("carte graphique", 30, 8, termbox.ColorRed)
		if p.IsInInventory("souris") {
			TermPrint("souris", 48, 8, termbox.ColorGreen)
		} else {
			TermPrint("souris", 48, 8, termbox.ColorRed)
		}
	}

	if p.IsInInventory("febreze") && p.inventaire["febreze"] >= 2 {
		TermPrint("2x febreze", 30, 9, termbox.ColorGreen)
		if p.IsInInventory("clavier mecanique") {
			TermPrint("clavier mecanique", 48, 9, termbox.ColorGreen)
		} else {
			TermPrint("clavier mecanique", 48, 9, termbox.ColorRed)
		}
	} else {
		TermPrint("2x febreze", 30, 9, termbox.ColorRed)
		if p.IsInInventory("clavier mecanique") {
			TermPrint("clavier mecanique", 48, 9, termbox.ColorGreen)
		} else {
			TermPrint("clavier mecanique", 48, 9, termbox.ColorRed)
		}
	}

	if p.IsInInventory("febreze") {
		TermPrint("febreze", 30, 10, termbox.ColorGreen)
		if p.IsInInventory("souris") {
			TermPrint("souris", 48, 10, termbox.ColorGreen)
		} else {
			TermPrint("souris", 48, 16, termbox.ColorRed)
		}
	} else {
		TermPrint("febreze", 30, 10, termbox.ColorRed)
		if p.IsInInventory("souris") {
			TermPrint("souris", 48, 10, termbox.ColorGreen)
		} else {
			TermPrint("souris", 48, 10, termbox.ColorRed)
		}
	}
	fmt.Println("\n\n\n\n-------------------------------------")
	fmt.Print("❖ Que veux tu parmis tous ces objets\n\n\n\n")
	answer := Scan() //Le mec il rentre le numero auquel est attribué à son objet
	if answer == "0" {
		p.Menu()
	}
	i, err := strconv.Atoi(answer) //string en int
	i -= 1
	if err != nil {
		fmt.Println("Pas compris")
		Enter()
		p.Forgeron()
	}
	if p.wallet >= forgeron.inventaire[lst[i]] && p.LimitSpace() {
		switch answer {
		case "1":
			if p.IsInInventory("carte graphique") && p.IsInInventory("souris") {
				p.RemoveInventory("carte graphique")
				p.RemoveInventory("souris")
				p.AddInventory("casque gaming")
				p.wallet -= forgeron.inventaire[lst[i]]
				Enter()
				p.Menu()
			}
		case "2":
			if p.IsInInventory("febreze") && p.IsInInventory("clavier mecanique") {
				p.RemoveInventory("clavier mecanique")
				p.RemoveInventory("febreze")
				if !p.IsInInventory("febreze") {
					fmt.Println("Tu n'as pas les objets requis pour te fabriquer cet objet ou t'a plus de place")
					p.AddInventory("febreze")
					p.AddInventory("clavier mecanique")
					Enter()
					p.Menu()
				} else {
					p.RemoveInventory("febreze")
					p.AddInventory("pull ynov")
					fmt.Println("tu possèdes maintenant un pull ynov ")
					p.wallet -= forgeron.inventaire[lst[i]]
					Enter()
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
				Enter()
				p.Menu()
			}
		case "0":
			p.Menu()
		default:
			fmt.Println("je n'ai pas compris ta réponse, peux tu repeter ? ")
			Enter()
			p.Forgeron()
		}
	} else {
		fmt.Println("tu ne possèdes pas suffisament d'argent ou ton inventaire est plein. reviens plus tard . ")
		Enter()
		p.Menu()
	}
	fmt.Println("Tu n'as pas les objets requis pour te fabriquer cet objet ou t'a plus de place")
	Enter()
	p.Menu()
}

func (p *Personnage) Equip() {
	lst := TransvalseList(p.inventaire)
	fmt.Println("-----------------------")
	fmt.Println("Ton inventaire est composé de")
	for ind, val := range lst {
		fmt.Printf("๑ %d %s pour %d € \n", ind+1, val, p.inventaire[val]) //le ind il sert juste pour numeroter les items la clé est le nom de l'objet et val est le montant
	}
	fmt.Println("\n----------------------")
	fmt.Print("❖ Que veux tu equiper parmis tous ces objets\n\n\n\n")
	answer := Scan()
	i, err := strconv.Atoi(answer)
	i -= 1
	if err != nil {
		fmt.Println("Pas compris")
		Enter()
		p.Equip()
	}
	switch answer {
	case "1":
		p.armure.body = "pull ynov"
		p.RemoveInventory("pull ynov")
		fmt.Println(("Ton pull ynov est maintenant equipé "))
		p.notemax += 10
		p.note += 10
		Enter()
		p.Menu()
	case "2":
		p.armure.hand = "multiprise"
		p.RemoveInventory("multiprise")
		fmt.Println((" ta multiprise est maintenant equipée "))
		p.strengh += 10
		Enter()
		p.Menu()
	case "3":
		p.armure.head = "casque gaming"
		p.RemoveInventory("casque gaming")
		fmt.Println((" ton casque gaming est maintenant equipé "))
		p.intmax += 10
		p.energy += 10
		Enter()
		p.Menu()
	default:
		fmt.Println("je n'ai pas compris ta réponse, peux tu repeter ? ")
		Enter()
		p.Equip()
	}
}

func (p *Personnage) Garden() {
	if p.IsInInventory("graine") {
		min = time.Now().Minute()
		fmt.Print("Bienvenue dans la terrasse\nVoulez vous planter votre graine\n1-Oui\n2-Non\n\n\n\n")
		ans := Scan()
		if ans == "1" {
			if Len(jardin) == 0 {
				jardin["graine 1"] = 0
				fmt.Println("Vous avez ajouté une graine")
				Enter()
				p.Menu()
			}
			fmt.Println("Voici votre jardin")
			jardin["graine "+strconv.Itoa(Len(jardin)+1)] = 0
			for val, cle := range jardin {
				jardin[val] = cle + min
				fmt.Println(val, cle, "point de puissance")
			}
			fmt.Print("Veux tu recolter une graine\n1-Oui\n2-Non\n\n\n\n")
			ans = Scan()
			if ans == "1" {
				fmt.Print("Laquelle?\n\n\n\n")
				ans = Scan()
				p.strengh += float64(jardin[ans])
				delete(jardin, ans)
				fmt.Println("Ta force est à", p.strengh)
				Enter()
				p.Menu()
			} else if ans == "2" {
				fmt.Println("Dégage")
				Enter()
				p.Menu()
			} else {
				fmt.Println("Pas compris")
				Enter()
				p.Garden()
			}
		} else if ans == "2" && Len(jardin) > 0 {
			fmt.Println("Dégage")
			Enter()
			p.Menu()
		}
		fmt.Println("Pas compris")
		Enter()
		p.Garden()
	} else {
		fmt.Println("T'as pas de graine salaud")
		Enter()
		p.Menu()
	}
}
