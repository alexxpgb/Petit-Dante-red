package piscine

import "fmt"

type Personnage struct {
	name       string
	classe     string
	niveau     string
	notemax    int
	note       int
	inventaire map[string]int
	skills     []string
	wallet     int
}

func (p *Personnage) Init(name string, classe string) {
	p.name = name
	p.classe = classe
	p.niveau = "B1"
	p.notemax = 100
	p.note = 50
	p.inventaire = map[string]int{"sucette": 3, "totem": 1}
	p.skills = []string{"python"}
	p.wallet = 50
}
func (p *Personnage) Init1() {
	var answer string
	var answer2 string
	fmt.Println("quel est ton nom ?")
	fmt.Scan(&answer)
	if IsUpper(string(answer[0])) && IsLower(string(answer[1:])) {
		p.name = answer
	} else if IsAlpha(answer) {
		answer2 = ToUpper(string(answer[0]))
		answer2 = ToLower(answer[1:])
		p.name = answer2
	} else {
		fmt.Println("Ton nom n'est pas valide, met en un autre.")
		p.Init1()
	}
	p.classe = "info"
	p.niveau = "B1"
	p.notemax = 100
	p.note = 50
	p.inventaire = map[string]int{"sucette": 3, "totem": 1}
	p.skills = []string{"python"}
	p.wallet = 50
}

func (p *Personnage) class() {
	var answer string
	if p.niveau == "B3" {
		fmt.Println("Quelle est ta specialisation ?")
		fmt.Println("-------------------------------")
		fmt.Println("1-IA data  2-infra 3-cybersécurité 4- dev  ")
		fmt.Println("Saisie le numéro de ta spécialité")
		fmt.Scan(&answer)
		if answer == "1" {
			p.classe = "IA data"
		} else if answer == "2" {
			p.classe = "infra "
		} else if answer == "3" {
			p.classe = "cybersécurité"
		} else if answer == "4" {
			p.classe = "dev"
		}
	}
}
