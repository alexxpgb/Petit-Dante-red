package piscine

import (
	"fmt"
	"math/rand"
	"time"
)

var m Mentor = Mentor{"Eleve", 100, 100, 3, 4, 50, 20}

func (m *Mentor) MentorPattern(p *Personnage, i int) {
	if i%3 == 0 { // tous les 3 tour la force du mentor double
		p.note -= m.strengh * 2
	} else {
		p.note -= m.strengh //Mon perso prend des dégats aie
	}
	fmt.Printf("%s a attaqué %s de %v point de degat \nTu est maintenant à %v/%v\n\n", m.name, p.name, m.strengh, p.note, p.notemax)
	if !p.IsAlive() { //Je regarde s'il est vivant et s'il peut ressuciter
		p.Redouble()
	}

}

func (p *Personnage) CharTurn(m *Mentor) { //Le systeme de combat pour mon joueur
	fmt.Println("C'est maintenant à vous de jouer")
	fmt.Println("Menu:")
	fmt.Println("1 - Attaquer")
	fmt.Print("2 - Inventaire\n\n\n\n\n\n\n")
	a := Scan()
	switch a {
	case "1":
		fmt.Print("Avec quels skills veut tu attaquer?\n\n") //On attaque avec nos skills
		for i, c := range p.skills {
			fmt.Println(i+1, "-", c) // J'affiche les skills avec leurs indice+1 pour commancer à 1
		}
		fmt.Print("\n\n\n\n\n")
		b := Scan()
		switch b {
		case "1": // A revoir
			if p.energy-bos["python"] >= 0 {
				p.energy -= bos["python"]
				m.note -= int(p.strengh)
				fmt.Printf("%s a attaqué %s de %v point de degat \nIl est maintenant à %v/%v\n", p.name, m.name, p.strengh, m.note, m.notemax)
			} else {
				fmt.Println("Tu n'a plus assez d'energie")
			}

		case "2":
			if p.IsInSkill("go") {
				if rand.Float64() < 0.5 { //1 chance sur deux
					if p.energy-bos["go"] >= 0 {
						p.energy -= bos["go"]
						m.note -= int(p.strengh) * 2
						fmt.Printf("%s a attaqué %s de %v point de degat \nIl est maintenant à %v/%v\n", p.name, m.name, p.strengh*2, m.note, m.notemax)
					}
				} else {
					fmt.Println("Tu a rater ton coup")
				}
			}
		}
	case "2":
		p.AccessInventory(2) //2=J'utilise un item en plein combat
	default:
		fmt.Println("Veuillez choisir une option valide")
	}
}

func (m *Mentor) Training(p *Personnage) {

	var count int = 1
	for p.IsAlive() && m.note > 0 { //Tant qu'il y en a un en vie
		if p.initiative > m.initiative { //s'il a plus d'initiative que moi il commence
			time.Sleep(time.Second * 1) //Juste pour que ce soit plus lissible et pratique
			p.CharTurn(m)
			m.MentorPattern(p, count)
		} else {
			time.Sleep(time.Second * 1)
			m.MentorPattern(p, count)
			p.CharTurn(m)
		}
		count++
	}
	if p.IsAlive() { //A finir normalement il devrait gagner des trucs s'il gagne genre exp initiative et sous peut être même des objets
		fmt.Println("Votre échauffement est maintenant terminé, vous avez gagné") //Je pense qu'on va faire en sorte que il regagne ses pv vu que normalement c'est qu'un pnj
		p.exp += m.exp
		p.LevelUp()
		p.initiative += m.initiative
		p.wallet += m.wallet
		p.RandomObjects(0.5)
	} else {
		fmt.Println("Votre échauffement est maintenant terminé, vous avez perdu")
	}
	p.Menu()
}
