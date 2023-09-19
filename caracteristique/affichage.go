package piscine

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

var p Personnage

func ReadInputO() {
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
				if choix == 2 || choix == 3 {
					choix--
				} else {
					choix = 3
				}
			case term.KeyArrowDown:
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
					term.Close()
					p.Init()
				case 2:
					term.Close()
					var p2 Personnage
					p2.Init()
					p.name = p2.name
				case 3:
					return
				}
			}
		case term.EventError:
			panic(ev.Err)
		}
		Graphisme(choix)
	}
	return
}

func Graphisme(choix int) {
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
		fmt.Println("▓                                                         • ▌ ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄      ▓")
		fmt.Println("▓                     ╔╗╔┌─┐┬ ┬  ┌─┐┌─┐┌┬┐┌─┐             ·██ ▐███▪▀▄.▀·•█▌▐█•██  ▪     ▀▄ █·    ▐▄▄· █ ▐█ ▀ ▪██▪▐█•██  ▀▄.▀·▀▄ █·    ▓")
		fmt.Println("▓                  1• ║║║├┤ │││  │ ┬├─┤│││├┤              ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█▄██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                     ╝╚╝└─┘└┴┘  └─┘┴ ┴┴ ┴└─┘             ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐███▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                        ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀     ▓")
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
		fmt.Println("▓                                                         • ▌  ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄      ▓")
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
		fmt.Println("▓                                                         • ▌  ▄ ·. ▄▄▄ . ▐ ▄ ▄▄▄▄▄      ▄▄▄      ·▄▄▄▪   ▄▄ •  ▄ .▄▄▄▄▄▄▄▄▄ .▄▄▄      ▓")
		fmt.Println("▓                                                         ·██ ▐███▪▀▄.▀·•█▌▐█•██  ▪     ▀▄ █·    ▐▄▄· █ ▐█ ▀ ▪██▪▐█•██  ▀▄.▀·▀▄ █·    ▓")
		fmt.Println("▓                            1• NEW GAME                  ▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌ ▐█.▪ ▄█▀▄ ▐▀▀▄     ██▪ ▐█·▄█ ▀█▄██▀▐█ ▐█.▪▐▀▀▪▄▐▀▀▄     ▓")
		fmt.Println("▓                                                         ██ ██▌▐█▌▐█▄▄▌██▐█▌ ▐█▌·▐█▌.▐▌▐█•█▌    █▌ .▐█▌▐█▄▪▐███▌▐▀ ▐█▌·▐█▄▄▌▐█•█▌    ▓")
		fmt.Println("▓                                                         ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀  ▀█▄▀▪.▀  ▀    ▀▀▀ ▀▀▀·▀▀▀▀ ▀▀▀ · ▀▀▀  ▀▀▀ .▀  ▀    ▓")
		fmt.Println("▓                            2• SETTINGS                                                                                              ▓")
		fmt.Println("▓                                                                                                                                     ▓")
		fmt.Println("▓                     ╔═╗ ┬ ┬┬┌┬┐                                                         __...--~~~~~-._   _.-~~~~~--...__           ▓")
		fmt.Println("▓                  3• ║═╬╗│ ││ │                                                        //               `V'              \\\\  	      ▓")
		fmt.Println("▓                     ╚═╝╚└─┘┴ ┴                                                    	//                 |                \\\\         ▓")
		fmt.Println("▓                                                                                     //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\       ▓")
		fmt.Println("▓                                                                                   //__.....----~~~~._\\\\ | //_.~~~~----.....__\\\\     ▓")
		fmt.Println("▓                                                                         	    ====================\\\\|//====================     ▓")
		fmt.Println("▓                                                                                           		 `---`                        ▓")
		fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")

	}
}

func (p *Personnage) Menu() {
	var answer int
	fmt.Println("Pour acceder à ton inventaire, tape 1. \nPour acceder aux informartions de ton personnage, tape 2. \nPour acceder à la peda , tape 3. \nPour acceder au forgeron, tape 4. \nPour retrouver au menu d'accueil, tape 5.")
	fmt.Scan(&answer)
	switch answer {
	case 1:
		p.AccessInventory()
	case 2:
		p.Display()

	case 3:
		p.Boutique()

	case 4:
		p.Forgeron()
	case 5:
		ReadInputO()

	default:
		fmt.Println("❖Je n'ai pas compris ta requête, peux tu repeter ? ")
		p.Menu()
	}
}

func (p Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("๑ Ton nom est :", p.name)
	fmt.Println("๑ Ta spécialité est :", p.classe)
	fmt.Println("๑ Ton niveau est :", p.niveau)
	fmt.Printf("๑ Tu as %d/%d\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %d %s\\n", val, cle)
	}
	fmt.Println("Ta liste de skills est :")
	for _, val := range p.skills {
		fmt.Println("๑ ", val)
	}
	fmt.Printf("๑ Tu as %d euros\\n", p.wallet)
	fmt.Println("-----------------------")
	p.Menu()
}
func (p Personnage) AccessInventory() { // ca permet d'accéder a ton inventaire batard
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf("๑ %d %s\n", val, cle)
	}
	fmt.Println("----------------------")
	fmt.Println("❖Veut tu utiliser un de ses objets?")
	fmt.Println("1- Oui")
	fmt.Println("2- Non")
	var answer int
	fmt.Scanln(&answer)
	switch answer {
	case 1:
		fmt.Println("❖Lequel?")
		var answer2 string
		fmt.Scanln(&answer2)
		p.UseObject(answer2)
	case 2:
		fmt.Println("Alors tu peux continuer")
		p.Menu()
	default:
		fmt.Println("❖Je n'ai pas compris ta requête, peux tu repeter? ")
		p.AccessInventory()
	}
}

func (p *Personnage) BookOfSkills(s string) { // fct qui permet d'apprendre des compétences
	bos := Personnage{name: "book", classe: "book", inventaire: map[string]int{"python": 1}}
	bos.inventaire[s] = 1
	fmt.Println("-----------------------")
	fmt.Println("❖Quels skills veut tu apprendre?")
	for cle, val := range bos.inventaire {
		fmt.Printf("๑ %d %s", val, cle)
	}
	fmt.Println("\\n----------------------")
	var answer string
	fmt.Scan(&answer)
	if !p.IsInSkill(answer) { // si je n'ai pas ce skill dans ma skill liste :
		if bos.IsInInventory(answer) { // et s'il est proposé dans l'inventaire :
			p.skills = append(p.skills, answer)      // je l'ajoute a ma liste de skills
			fmt.Println("Vous avez appris ce skill") // la si t'as pas compris t'es un sale PD
		} else {
			fmt.Println("Le skill que tu m'a proposé ne fait pas partie de ma liste de skills")
		}
	} else {
		fmt.Println("Vous avez déjà appris ce skill")
	}
}
