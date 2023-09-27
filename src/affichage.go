package piscine

import (
	"fmt"
	"strconv"

	"github.com/nsf/termbox-go"
	term "github.com/nsf/termbox-go" //En vrai sans ça le jeu il serai fini en 2jour donc tant mieux que je l'ai mit pour me faire chier
)

var p Personnage
var bos map[string]int = map[string]int{"python": 5}
var check bool
var color termbox.Attribute = termbox.ColorDefault

func ReadInputO() { //Pour pouvoir faire le mouvement des touches flèche hauts et bas
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	inputchek := false
	var choix int = 1
	Graphisme(choix)
	for !inputchek {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowUp:
				fmt.Print("\033[H\033[2J")

				if choix == 2 || choix == 3 {
					choix--
				} else {
					choix = 3
				}
			case term.KeyArrowDown:
				fmt.Print("\033[H\033[2J") //Clear terminal
				if choix == 1 || choix == 2 {
					choix++
				} else {
					choix = 1
				}
			case term.KeyEnter:
				fmt.Print("\033[H\033[2J")
				inputchek = true
				switch choix {
				case 1:
					if check {
						p.Menu()
					}
					p.Init()
				case 2:
					fmt.Println("Voulez-vous changer la couleur du jeu\n1-Oui\n2-Non")
					a := Scan()
					if a == "1" {
						fmt.Println("Quel couleur voulez-vous\n1-Rouge\n2-Bleu\n3-Vert")
						b := Scan()
						if b == "1" {
							color = termbox.ColorRed
						} else if b == "2" {
							color = termbox.ColorCyan
						} else if b == "3" {
							color = termbox.ColorGreen
						} else {
							fmt.Println("Il était pas si compliqué ce choix")
						}
					}
					fmt.Println("Voulez vous recommencer le jeux\n1-Oui\n2-Non")
					a = Scan()
					if a == "1" {
						fmt.Print("\033[H\033[2J")
						var p2 Personnage
						p2.Init()
					} else if a == "2" {
						ReadInputO()
					} else {
						fmt.Println("Il était pas si compliqué ce choix")
						ReadInputO()
					}
				case 3:
					return
				}
			}
		case term.EventError:
			panic(ev.Err)
		}
		Graphisme(choix)
	}
	fmt.Print("\033[H\033[2J")
}

func Graphisme(choix int) { //Mes affichage  de menu principal
	switch choix {
	case 1:
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓  _       __     __                             __                                                                                   ▓")
		fmt.Println("▓  | |     / ___  / _________  ____ ___  ___     / /_____                                                                             ▓")
		fmt.Println("▓  | | /| / / _ \\/ / ___/ __ \\/ __ `__ \\/ _ \\   / __/ __ \\                                                                            ▓")
		fmt.Println("▓  | |/ |/ /  __/ / /__/ /_/ / / / / / /  __/  / /_/ /_/ /                                                                            ▓")
		fmt.Println("▓  |__/|__/\\___/_/\\___/\\____/_/ /_/ /_/\\___/   \\__/\\____/                                                                             ▓")
		fmt.Println("▓                             			                                                                                      ▓")
		fmt.Println("▓                                                         • ▌  ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄     ▓")
		fmt.Println("▓                     ╔╗╔┌─┐┬ ┬  ┌─┐┌─┐┌┬┐┌─┐             ·██ ▐███▪▀▄.▀·•█▌▐█•██  ▪     ▀▄ █·    ▐▄▄· █ ▐█ ▀ ▪██▪▐█•██  ▀▄.▀·▀▄ █·    ▓")
		fmt.Println("▓                   • ║║║├┤ │││  │ ┬├─┤│││├┤              ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█ ██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                     ╝╚╝└─┘└┴┘  └─┘┴ ┴┴ ┴└─┘             ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐ ██▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                         ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                              • SETTINGS                                                                                             ▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓                                                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                              • QUIT                                                   //               `V'              \\\\  	      ▓")
		fmt.Println("▓                                                                                      //                 |                \\\\         ▓")
		fmt.Println("▓                                                                                     //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\       ▓")
		fmt.Println("▓                                                                                   //__.....----~~~~._\\\\ | //_.~~~~----.....__\\\\     ▓")
		fmt.Println("▓                                                                         	    ====================\\\\|//====================     ▓")
		fmt.Println("▓                                                                                           		 `---`                        ▓")
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	case 2:
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓  _       __     __                             __                                                                                   ▓")
		fmt.Println("▓  | |     / ___  / _________  ____ ___  ___     / /_____                                                                             ▓")
		fmt.Println("▓  | | /| / / _ \\/ / ___/ __ \\/ __ `__ \\/ _ \\   / __/ __ \\                                                                            ▓")
		fmt.Println("▓  | |/ |/ /  __/ / /__/ /_/ / / / / / /  __/  / /_/ /_/ /                                                                            ▓")
		fmt.Println("▓  |__/|__/\\___/_/\\___/\\____/_/ /_/ /_/\\___/   \\__/\\____/                                                                             ▓")
		fmt.Println("▓                             			                                                                                      ▓")
		fmt.Println("▓                                                         • ▌  ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄     ▓")
		fmt.Println("▓                                                         ·██ ▐███▪▀▄.▀·•█▌▐█•██  ▪     ▀▄ █·    ▐▄▄· █ ▐█ ▀ ▪██▪▐█•██  ▀▄.▀·▀▄ █·    ▓")
		fmt.Println("▓                             • NEW GAME                  ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█ ██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                                                         ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐ ██▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                     ╔═╗┌─┐┌┬┐┌┬┐┬┌┐┌┌─┐┌─┐              ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                   • ╚═╗├┤  │  │ │││││ ┬└─┐                                                                                          ▓")
		fmt.Println("▓                     ╚═╝└─┘ ┴  ┴ ┴┘└┘└─┘└─┘                                                                                          ▓")
		fmt.Println("▓                                                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                              • QUIT                                                   //               `V'              \\\\  	      ▓")
		fmt.Println("▓                                                                                      //                 |                \\\\         ▓")
		fmt.Println("▓                                                                                     //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\       ▓")
		fmt.Println("▓                                                                                   //__.....----~~~~._\\\\ | //_.~~~~----.....__\\\\     ▓")
		fmt.Println("▓                                                                         	    ====================\\\\|//====================     ▓")
		fmt.Println("▓                                                                                           		 `---`                        ▓")
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	case 3:
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓  _       __     __                             __                                                                                   ▓")
		fmt.Println("▓  | |     / ___  / _________  ____ ___  ___     / /_____                                                                             ▓")
		fmt.Println("▓  | | /| / / _ \\/ / ___/ __ \\/ __ `__ \\/ _ \\   / __/ __ \\                                                                            ▓")
		fmt.Println("▓  | |/ |/ /  __/ / /__/ /_/ / / / / / /  __/  / /_/ /_/ /                                                                            ▓")
		fmt.Println("▓  |__/|__/\\___/_/\\___/\\____/_/ /_/ /_/\\___/   \\__/\\____/                                                                             ▓")
		fmt.Println("▓                             			                                                                                      ▓")
		fmt.Println("▓                                                         • ▌  ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄     ▓")
		fmt.Println("▓                                                         ·██ ▐███▪▀▄.▀·•█▌▐█•██  ▪     ▀▄ █·    ▐▄▄· █ ▐█ ▀ ▪██▪▐█•██  ▀▄.▀·▀▄ █·    ▓")
		fmt.Println("▓                             • NEW GAME                  ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█ ██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                                                         ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐ ██▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                         ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                             • SETTINGS                                                                                              ▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓                     ╔═╗ ┬ ┬┬┌┬┐                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                   • ║═╬╗│ ││ │                                                        //               `V'              \\\\  	      ▓")
		fmt.Println("▓                     ╚═╝╚└─┘┴ ┴                                                       //                 |                \\\\         ▓")
		fmt.Println("▓                                                                                     //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\       ▓")
		fmt.Println("▓                                                                                   //__.....----~~~~._\\\\ | //_.~~~~----.....__\\\\     ▓")
		fmt.Println("▓                                                                         	    ====================\\\\|//====================     ▓")
		fmt.Println("▓                                                                                           		 `---`                        ▓")
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	}
}

func (p *Personnage) Menu() {
	fmt.Print("\033[H\033[2J")
	check = true
	fmt.Print("\n\n\n\n\n\n")
	TermPrint(" _ __ ___   ___ _ __  _   _ ", 1, 1, termbox.ColorCyan)
	TermPrint("| '_ ` _ \\ / _ \\ '_ \\| | | |", 1, 2, termbox.ColorCyan)
	TermPrint("| | | | | |  __/ | | | |_| |", 1, 3, termbox.ColorCyan)
	TermPrint("|_| |_| |_|\\___|_| |_|\\__,_|", 1, 4, termbox.ColorCyan)
	fmt.Println("--------------------------------------------------------")
	fmt.Println("1/Pour acceder à ton inventaire. \n2/Pour acceder aux informartions de ton personnage. \n3/Pour acceder à la peda . \n4/Pour acceder a l'admin. \n5/Pour acceder a la liste de skill dans ta bibliothèque \n6/Pour aller s'entrainer \n7/Pour commencer le mode histoire\n8/Qui sont t-ils? \n9/Pour revenir au menu principal\n0/Pour allez au terrasse tapez 0\n Pour utiliser un ou des equipements, tape e")
	fmt.Println("--------------------------------------------------------")
	switch ev := term.PollEvent(); ev.Type {
	case term.EventKey: //Avec sa a peine tu touche une touche instant la demande est envoyé (pas besoin d'appuyer sur entrée)
		switch ev.Ch {
		case '1':
			p.AccessInventory(1)
		case '2':
			p.Display()

		case '3':
			p.Boutique()

		case '4':
			p.Forgeron()
		case '5':
			p.BookOfSkills()
		case '6':
			m.Training(p)
		case '7':
			fmt.Println("En construction")
			Enter()
			p.Menu()
		case '8':
			fmt.Println("ABBA ,Spielberg et QUEEN")
			Enter()
			p.Menu()
		case '9':
			ReadInputO()
		case 'e':
			p.Equip()
		case '0':
			p.Garden()
		default:
			fmt.Println("❖ Je n'ai pas compris ta requête, peux tu repeter ? ")
			p.Menu()
		}
	}
}
func (p *Personnage) Display() { //A modifier
	var chck bool
	fmt.Println("-----------------------")
	fmt.Println("๑ Ton nom est :", p.name)
	fmt.Println("๑ Ta spécialité est :", p.classe)
	fmt.Println("๑ Ton niveau est :", p.niveau)
	fmt.Printf("๑ Tu as %v/%v\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %v %s\n", val, cle)
	}
	fmt.Println("Ta liste de skills est :")
	for _, val := range p.skills {
		fmt.Println(" • ", val)
	}
	fmt.Printf("๑ Tu as %v euros\n", p.wallet)
	fmt.Printf("๑ Tu as %v point de force\n", p.strengh)
	fmt.Printf("๑ Tu as %v/%v force mental\n", p.energy, p.intmax)
	fmt.Printf("๑ Tu as %v point d'initiative\n", p.initiative)
	fmt.Printf("๑ Tu as %v/%v point d'experience\n", p.exp, p.expmax)
	fmt.Println("-----------------------")
	for !chck {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEnter:
				chck = true
			}
		}
	}
	p.Menu()
}
func (p *Personnage) AccessInventory(nb int) { // ca permet d'accéder a ton inventaire
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %d %s\n", val, cle)
	}
	fmt.Println("----------------------")
	fmt.Println("❖ Veut tu utiliser un de ses objets?")
	fmt.Println("1- Oui")
	fmt.Print("2- Non\n\n\n\n\n\n") //Pour qu'on puisse voir l'endroit des réponses
	switch Scan() {
	case "1":
		fmt.Print("❖Lequel?\n\n\n\n\n\n")
		ans := Scan()
		p.UseObject(&m, ans, nb)
		Enter()
		if nb == 1 {
			p.Menu()
		}
	case "2":
		fmt.Println("Alors tu peux continuer")
		if nb == 1 {
			p.Menu()
		}
	default:
		fmt.Println("❖ Je n'ai pas compris ta requête, peux tu repeter? ")
		p.AccessInventory(nb)
	}
}

func AppendSkill(s string, m int) {
	if IsInMap(bos, s) {
		fmt.Println("Tu as déjà ce skill")
	} else {
		bos[s] = m
	}
}

func (p *Personnage) BookOfSkills() { // fct qui permet d'apprendre des compétences
	lst := TransvalseList(bos) //Je cast ma map en liste
	fmt.Println("-----------------------")
	fmt.Println("❖ Quels skills veut tu apprendre?")
	for i, skl := range lst {
		fmt.Printf(" • %d %s (Off: %d Int)\n", i+1, skl, bos[skl])
	}
	fmt.Print("\n----------------------\n\n\n\n")
	answer := Scan()                   //Sans doute une erreur par la
	answint, _ := strconv.Atoi(answer) //string en int
	answint -= 1
	if !p.IsInSkill(answer) { // si je n'ai pas ce skill dans ma skill liste :
		if answint < len(lst) && answint > 0 { // et s'il est proposé dans l'inventaire :
			p.skills = append(p.skills, lst[answint]) // je l'ajoute a ma liste de skills
			fmt.Println("Vous avez appris ce skill")
			Enter()
			p.Menu()
		} else {
			fmt.Println("Le skill que tu m'a proposé ne fait pas partie de ma liste de skills")
			Enter()
			p.Menu()
		}
	} else {
		fmt.Println("Vous avez déjà appris ce skill")
		Enter()
		p.Menu()
	}
}
