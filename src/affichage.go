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
	fmt.Println("1/Pour accéder à ton inventaire. \n2/Pour accéder aux informations de ton personnage. \n3/Pour accéder à la PEDA . \n4/Pour accéder a l'admin. \n5/Pour accéder à la liste de skill dans ta bibliothèque \n6/Pour aller s'entrainer \n7/Pour commencer le mode histoire\n8/Pour aller au terrasse \n9/Pour revenir au menu principal\n0/Qui sont t-ils?")
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
			p.Story()
		case '8':
			p.Garden()
		case '9':
			ReadInputO()
		case '0':
			fmt.Println("ABBA ,Spielberg et QUEEN")
			Enter()
			p.Menu()
		default:
			fmt.Println("❖ Je n'ai pas compris ta requête, peux tu répèter ? ")
			p.Menu()
		}
	}
}
func (p *Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("๑ Ton nom est :", p.name)
	fmt.Println("๑ Ta spécialité est :", p.classe)
	fmt.Println("๑ Ton niveau est :", p.niveau)
	fmt.Printf("๑ Ta note est de %v/%v\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("• %v %s\n", val, cle)
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
	fmt.Println("  __")
	fmt.Println(" (๏๏)←", p.armure.head)
	fmt.Printf(" /||\\←%v\n", p.armure.body)
	fmt.Println("° /\\ ° ←", p.armure.hand)

	fmt.Println("-----------------------")
	Enter()
	p.Menu()
}
func (p *Personnage) AccessInventory(nb int) { // ca permet d'accéder a ton inventaire
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %d %s\n", val, cle)
	}
	fmt.Println("----------------------")
	fmt.Println("❖ Veux-tu utiliser un de ces objets?")
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
		fmt.Printf(" •%d/ %s (Off: %d Int)\n", i+1, skl, bos[skl])
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

func (p *Personnage) Story() {
	if p.niveau != "B1" {
		Printtime("Il était une fois un jeune étudiant nommé ")
		fmt.Print(p.name)
		Printtime(", qui avait été accepté à l'école Ynov, réputée pour former les esprits les plus brillants de la nouvelle génération.")
		fmt.Println("")
		fmt.Print(p.name)
		Printtime(" était excité à l'idée de commencer sa formation dans cette prestigieuse institution, mais il était loin de se douter des défis qui l'attendaient.")
		fmt.Println("")
		Printtime("Dès son premier jour à Ynov, ")
		fmt.Print(p.name)
		Printtime("fut accueilli par les mentors, des professeurs réputés pour leur exigence et leur rigueur.")
		fmt.Println("")
		Printtime("Cependant, il ne tarda pas à découvrir que ces mentors étaient bien plus que de simples enseignants.")
		fmt.Println("")
		Printtime("Ils semblaient posséder des pouvoirs mystérieux et étaient déterminés à mettre à l'épreuve les compétences et la détermination de leurs étudiants.")
		fmt.Println("")
		fmt.Print(p.name)
		Printtime(" avait entendu parler des mentors avant de rejoindre l'école, mais il ne s'attendait pas à ce qu'ils soient si redoutables.")
		Printtime(" Chacun d'entre eux avait une quête particulière à accomplir, une épreuve à surmonter, avant de pouvoir prétendre à l'obtention de son diplôme.")
		Printtime(" Ces quêtes étaient réputées pour être extrêmement difficiles, et peu d'étudiants les avaient réussies par le passé")
		fmt.Println("")
		fmt.Println("")
		Printtime("4 mentors viennent Alan ,Ethan ,Kheir ,et Cyril un combat va arriver")
		fmt.Println("")
		Printtime("Choissisez votre adversaire")
		fmt.Println("")
		Printtime("1/Alan")
		fmt.Println("")
		Printtime("2/Ethan")
		fmt.Println("")
		Printtime("3/Kheir")
		fmt.Println("")
		Printtime("4/Cyril")
		fmt.Println("")
		var Alan Mentor = Mentor{"Alan", 285, 142, 24, 500, 35, 500}
		var Ethan Mentor = Mentor{"Ethan", 285, 160, 24, 300, 435, 500}
		var Kheir Mentor = Mentor{"Kheir", 170, 85, 20, 150, 7, 250}
		var Cyril Mentor = Mentor{"Cyril", 170, 85, 15, 150, 20, 250}
		answer := Scan()
		switch answer {
		case "1":
			Alan.Training(p)
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" rencontra le mentor Allan.")
			fmt.Println("")
			Printtime("Il était un expert en dev d’appli et défia ")
			fmt.Print(p.name)
			Printtime(" de maîtriser les techniques lunar et de capacitor les plus avancées. ")
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" fut envoyé dans une forêt mystérieuse où il dut affronter des créatures légendaires et maîtriser les mouvements complexes du my sql .")
			fmt.Println("")
			Printtime(" Il était déterminé à ne pas échouer, car il savait que son avenir à Ynov en dépendait.")
		case "2":
			Ethan.Training(p)
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" rencontra le mentor Ethan,il était un génie en dev gaming.")
			fmt.Println("")
			Printtime(" Il posa à ")
			fmt.Print(p.name)
			Printtime(" un défi intellectuel qui le plongea dans un labyrinthe de problèmes mathématiques complexes et d'expériences scientifiques dangereuses.")
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" passa des nuits blanches à résoudre des énigmes, mais il ne se laissa pas décourager.")
		case "3":
			Kheir.Training(p)
			Printtime("Lucas rencontra le mentor Kheir-Eddine, il était un philosophe renommé qui enseignait la sagesse et la réflexion profonde. Il guida Lucas dans une quête intérieure pour explorer ses croyances et ses valeurs, l'aidant à trouver un sens plus profond dans son parcours éducatif. Sous la direction de Kheir-Eddine, Lucas apprit à réfléchir de manière critique et à prendre des décisions éclairées.")
		case "4":
			Cyril.Training(p)
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" rencontra le mystérieux Cyril, un artiste renommé.")
			fmt.Println("")
			Printtime(" Il demanda à ")
			fmt.Print(p.name)
			Printtime(" de créer une œuvre d'art qui refléterait son âme.")
			fmt.Println("")
			fmt.Print(p.name)
			Printtime(" se mit à travailler sur une toile géante, exprimant ses émotions les plus profondes à travers la peinture.")
			fmt.Println("")
			Printtime(" Il découvrit que l'art était bien plus qu'une simple technique, c'était une façon de communiquer avec son moi intérieur.")

		default:
			fmt.Println("No comprendo")
		}
		Printtime("Lors de la cérémonie de remise des diplômes, les mentors lui remirent son parchemin avec fierté.")
		Printtime(" Ils avaient vu en ")
		fmt.Print(p.name)
		Printtime(" le potentiel qui sommeillait en lui et l'avaient aidé à le réaliser.")
		fmt.Println("")
		fmt.Print(p.name)
		Printtime(" avait non seulement obtenu son diplôme, mais il avait aussi gagné le respect et l'amitié des mentors qui l'avaient autrefois intimidé.")
		Printtime("Cette histoire de ")
		fmt.Print(p.name)
		Printtime(" à l'école Ynov rappelle que parfois, nos plus grands défis sont en réalité nos plus grandes opportunités.")
		Printtime("Il avait combattu ses \"bourreaux\" mentors, mais en réalité, il avait combattu ses propres doutes et ses limites pour devenir la meilleure version de lui-même.")
		fmt.Println("")
		fmt.Println("_______  _______  _______  _______    _______           _______  _______ ")
		fmt.Println("(  ____ \\(  ___  )(       )(  ____ \\  (  ___  )|\\     /|(  ____ \\(  ____ )		")
		fmt.Println("| (    \\/| (   ) || () () || (    \\/  | (   ) || )   ( || (    \\/| (    )|")
		fmt.Println("| |      | (___) || || || || (__      | |   | || |   | || (__    | (____)|")
		fmt.Println("| | ____ |  ___  || |(_)| ||  __)     | |   | |( (   ) )|  __)   |     __)")
		fmt.Println("| | \\_  )| (   ) || |   | || (        | |   | | \\ \\_/ / | (      | (\\ (   ")
		fmt.Println("| (___) || )   ( || )   ( || (____/\\  | (___) |  \\   /  | (____/\\| ) \\ \\__")
		fmt.Println("(_______)|/     \\||/     \\|(_______/  (_______)   \\_/   (_______/|/   \\__/")

	} else {
		fmt.Println("Le mode histoire n'est que disponible au niveau B2")
		Enter()
		p.Menu()
	}
}
