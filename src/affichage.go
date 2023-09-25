package piscine

import (
	"fmt"

	term "github.com/nsf/termbox-go" //En vrai sans ça le jeu il serai fini en 2jour donc tant mieux que je l'ai mit pour me faire chier
)

var p Personnage

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
					p.Init()
				case 2:
					fmt.Println("Voulez vous recommencer le jeux\n1-Oui\n2-Non")
					a := Scan()
					if a == "1" {
						var p2 Personnage
						p2.Init()
					} else if a == "2" {
						p.Menu()
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
		fmt.Println("▓                  1• ║║║├┤ │││  │ ┬├─┤│││├┤              ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█▄██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                     ╝╚╝└─┘└┴┘  └─┘┴ ┴┴ ┴└─┘             ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐███▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                         ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                             2• SETTINGS                                                                                             ▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓                                                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                             3• QUIT                                                   //               `V'              \\\\  	      ▓")
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
		fmt.Println("▓                            1• NEW GAME                  ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█▄██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                                                         ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐███▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                     ╔═╗┌─┐┌┬┐┌┬┐┬┌┐┌┌─┐┌─┐              ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                  2• ╚═╗├┤  │  │ │││││ ┬└─┐                                                                                          ▓")
		fmt.Println("▓                     ╚═╝└─┘ ┴  ┴ ┴┘└┘└─┘└─┘                                                                                          ▓")
		fmt.Println("▓                                                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                             3• QUIT                                                   //               `V'              \\\\  	      ▓")
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
		fmt.Println("▓                            1• NEW GAME                  ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█▄██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                                                         ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐███▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                         ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                            2• SETTINGS                                                                                              ▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓                     ╔═╗ ┬ ┬┬┌┬┐                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                  3• ║═╬╗│ ││ │                                                        //               `V'              \\\\  	      ▓")
		fmt.Println("▓                     ╚═╝╚└─┘┴ ┴                                                       //                 |                \\\\         ▓")
		fmt.Println("▓                                                                                     //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\       ▓")
		fmt.Println("▓                                                                                   //__.....----~~~~._\\\\ | //_.~~~~----.....__\\\\     ▓")
		fmt.Println("▓                                                                         	    ====================\\\\|//====================     ▓")
		fmt.Println("▓                                                                                           		 `---`                        ▓")
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	}
}

func (p *Personnage) Menu() {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Pour acceder à ton inventaire, tape 1. \nPour acceder aux informartions de ton personnage, tape 2. \nPour acceder à la peda , tape 3. \nPour acceder au forgeron, tape 4. \nPour acceder a la liste de skill dans ta bibliothèque tape 5 \nPour aller s'entrainer tape 6 \nPour commencer le mode histoire tape 7\nQui sont t-ils? tape 8 \nPour revenir au menu principal tape 9 ")
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
			p.Menu()
		case '8':
			fmt.Println("ABBA ,Spielberg et QUEEN")
			p.Menu()
		case '9':
			ReadInputO()
		default:
			fmt.Println("❖ Je n'ai pas compris ta requête, peux tu repeter ? ")
			p.Menu()
		}
	}
}
func (p *Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("๑ Ton nom est :", p.name)
	fmt.Println("๑ Ta spécialité est :", p.classe)
	fmt.Println("๑ Ton niveau est :", p.niveau)
	fmt.Printf("๑ Tu as %d/%d\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %d %s\n", val, cle)
	}
	fmt.Println("Ta liste de skills est :")
	for _, val := range p.skills {
		fmt.Println(" • ", val)
	}
	fmt.Printf("๑ Tu as %d euros\n", p.wallet)
	fmt.Println("-----------------------")
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
	fmt.Println("2- Non\n\n\n\n\n") //Pour qu'on puisse voir l'endroit des réponses
	switch Scan() {
	case "1":
		fmt.Println("❖Lequel?\n\n\n\n\n")
		ans := Scan()
		p.UseObject(m, ans, nb)
		if nb == 1 {
			p.Menu()
		}
	case "2":
		fmt.Println("Alors tu peux continuer")
		if nb == 1 {
			p.Menu()
		}
	default:
		fmt.Println("❖Je n'ai pas compris ta requête, peux tu repeter? ")
		p.AccessInventory(nb)
	}
}

func (p *Personnage) AppendSkill(s string) {
	if bos.inventaire[s] == 1 {
		fmt.Println("Tu as déjà ce skill")
	} else {
		bos.inventaire[s] = 1
	}
}

func (p *Personnage) BookOfSkills() { // fct qui permet d'apprendre des compétences
	bos := Personnage{name: "book", classe: "book", inventaire: map[string]int{"python": 1}}
	fmt.Println("-----------------------")
	fmt.Println("❖ Quels skills veut tu apprendre?")
	for cle, val := range bos.inventaire {
		fmt.Printf("๑ %d %s", val, cle)
	}
	fmt.Println("\n----------------------")
	answer := Scan()          //Sans doute une erreur par la
	if !p.IsInSkill(answer) { // si je n'ai pas ce skill dans ma skill liste :
		if bos.IsInInventory(answer) { // et s'il est proposé dans l'inventaire :
			p.skills = append(p.skills, answer) // je l'ajoute a ma liste de skills
			fmt.Println("Vous avez appris ce skill")
		} else {
			fmt.Println("Le skill que tu m'a proposé ne fait pas partie de ma liste de skills")
		}
	} else {
		fmt.Println("Vous avez déjà appris ce skill")
	}
}
