package piscine

import "fmt"

func Graphisme() {
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("▓                                                                                                                                     ▓")
	fmt.Println("▓  _       __     __                             __           __  ___           __                _______       __    __              ▓")
	fmt.Println("▓  | |     / ___  / _________  ____ ___  ___     / /_____     /  |/  ___  ____  / /_____  _____   / ____(_____ _/ /_  / /____  _____  ▓")
	fmt.Println("▓  | | /| / / _ \\/ / ___/ __ \\/ __ `__ \\/ _ \\   / __/ __ \\   / /|_/ / _ \\/ __ \\/ __/ __ \\/ ___/  / /_  / / __ `/ __ \\/ __/ _ \\/ ___/  ▓")
	fmt.Println("▓  | |/ |/ /  __/ / /__/ /_/ / / / / / /  __/  / /_/ /_/ /  / /  / /  __/ / / / /_/ /_/ / /     / __/ / / /_/ / / / / /_/  __/ /      ▓")
	fmt.Println("▓  |__/|__/\\___/_/\\___/\\____/_/ /_/ /_/\\___/   \\__/\\____/  /_/  /_/\\___/_/ /_/\\__/\\____/_/     /_/   /_/\\__, /_/ /_/\\__/\\___/_/       ▓")
	fmt.Println("▓                             			                                                                                    		  ▓")
	fmt.Println("▓                                                                                                                            		  ▓")
	fmt.Println("▓                                                                                                                            	      ▓")
	fmt.Println("▓                             1• NEW GAME                                                                                      	  ▓")
	fmt.Println("▓                                                                                                                            	      ▓")
	fmt.Println("▓                                                                                                                         		      ▓")
	fmt.Println("▓                             2• SETTINGS                                                                                		      ▓")
	fmt.Println("▓                                                                                                                         		      ▓")
	fmt.Println("▓                                                                                      __...--~~~~~-._   _.-~~~~~--...__      		  ▓")
	fmt.Println("▓                             3• QUIT                                                 //               `V'              \\\\  	      ▓")
	fmt.Println("▓                                                                                    //                 |                \\\\		   ▓")
	fmt.Println("▓                                                                                      //__...--~~~~~~-._  |  _.-~~~~~~--...__\\\\     ▓")
	fmt.Println("▓                                                                                    //__.....----~~~~._\\\\ | /_.~~~~----.....__\\\\    ▓")
	fmt.Println("▓                                                                         	     ====================\\\\|//====================   ▓")
	fmt.Println("▓                                                                                           		 `---`                       ▓")
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	var answer int
	fmt.Scanln(&answer)
	switch answer {
	case 1:
		fmt.Println("New Game")
		var p Personnage
		p.Init1()
		p.Menu()
	case 2:
		fmt.Println("en construction...")
	case 3:
		fmt.Println("Bye!")
		return
	default:
		fmt.Println("Wrong answer!")
	}
}
func (p *Personnage) Menu() {
	var answer int
	fmt.Println("pour acceder à ton inventaire, tape 1. Pour acceder aux informartions de ton personnage, tape 2. Enfin, pour acceder à la peda , tape 3  ")
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

	default:
		fmt.Println("Je n'ai pas compris ta requête, peux tu repeter ? ")

	}
}

func (p Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("Ton nom est :", p.name)
	fmt.Println("Ta spécialité est :", p.classe)
	fmt.Println("Ton niveau est :", p.niveau)
	fmt.Printf("Tu as %d/%d\\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("%d %s\\n", val, cle)
	}
	fmt.Println("Ta liste de skills est :")
	for _, val := range p.skills {
		fmt.Println(val)
	}
	fmt.Printf("Tu as %d euros\\n", p.wallet)
	fmt.Println("-----------------------")
}
func (p Personnage) AccessInventory() { // ca permet d'accéder a ton inventaire batard
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf(" %d %s", val, cle)
	}
	fmt.Println("\\n----------------------")
	fmt.Println("Veut tu utiliser un de ses objets?")
	fmt.Println("1- Oui")
	fmt.Println("2- Non")
	var answer int
	fmt.Scanln(&answer)
	switch answer {
	case 1:
		fmt.Println("Lequel?")
		var answer2 string
		fmt.Scanln(&answer2)
		p.UseObject(answer2)
	case 2:
		fmt.Println("Alors tu peux continuer")
	default:
		fmt.Println("Je n'ai pas compris ta requête, peux tu repeter? ")
		p.AccessInventory()
	}
}

func (p *Personnage) BookOfSkills(s string) { // fct qui permet d'apprendre des compétences
	bos := Personnage{name: "book", classe: "book", inventaire: map[string]int{"python": 1}}
	bos.inventaire[s] = 1
	fmt.Println("-----------------------")
	fmt.Println("Quels skills veut tu apprendre?")
	for cle, val := range bos.inventaire {
		fmt.Printf(" %d %s", val, cle)
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
