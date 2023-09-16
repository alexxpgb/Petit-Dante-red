package piscine

import "fmt"

func Graphisme() {
	fmt.Println("#######################################################################################")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                   WELCOME TO                                        #")
	fmt.Println("#                                 MENTOR FIGHTER                                      #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                             1- NEW GAME                                             #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                             2- SETTINGS                                             #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                             3- QUIT                                                 #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#                                                                                     #")
	fmt.Println("#######################################################################################")
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

	default:
		fmt.Println("Je n'ai pas compris ta requête, peux tu repeter ? ")

	}
}

func (p Personnage) Display() {
	fmt.Println("-----------------------")
	fmt.Println("Ton nom est :", p.name)
	fmt.Println("Ta spécialité est :", p.classe)
	fmt.Println("Ton niveau est :", p.niveau)
	fmt.Printf("Tu as %d/%d\n", p.note, p.notemax)
	fmt.Println("Dans ton inventaire tu as :")
	for cle, val := range p.inventaire {
		fmt.Printf("%d %s\n", val, cle)
	}
	fmt.Println("Ta liste de skills est :")
	for _, val := range p.skills {
		fmt.Println(val)
	}
	fmt.Printf("Tu as %d euros\n", p.wallet)
	fmt.Println("-----------------------")
}
func (p Personnage) AccessInventory() {
	fmt.Println("-----------------------")
	fmt.Println("L'inventaire est composé de")
	for cle, val := range p.inventaire {
		fmt.Printf(" %d %s", val, cle)
	}
	fmt.Println("\n----------------------")
}
