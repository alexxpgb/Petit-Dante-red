package piscine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var round int = 4
var m Mentor = Mentor{"Eleve", 100, 50, 3, 50, 4, 20}

func (m *Mentor) MentorPattern(p *Personnage, i int) {
	if i%3 == 0 { // tous les 3 tour la force du mentor double
		p.note -= m.strengh * 2
		Enter()
		fmt.Print("\033[H\033[2J")
		TermPrint("Damage Critique", 20, 0, termbox.ColorRed)
		fmt.Printf("\n%s a attaqué %s de %v point de degats \nTu est maintenant à %v/%v\n\n", m.name, p.name, m.strengh*2, p.note, p.notemax)
		if !p.IsAlive() { //Je regarde s'il est vivant et s'il peut ressuciter
			p.Redouble()
		}
	} else {
		p.note -= m.strengh //Mon perso prend des dégats aie
		fmt.Printf("%s a attaqué %s de %v point de degats \nTu est maintenant à %v/%v\n\n", m.name, p.name, m.strengh, p.note, p.notemax)
		if !p.IsAlive() { //Je regarde s'il est vivant et s'il peut ressuciter
			p.Redouble()
		}
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
		fmt.Print("Avec quels skills veux tu attaquer?\n\n") //On attaque avec nos skills
		fmt.Println("1 / Cuisine")
		for i, c := range p.skills {
			fmt.Println(i+2, "/", c) // J'affiche les skills avec leurs indice+1 pour commancer à 1
		}
		fmt.Print("\n\n\n\n\n")
		b := Scan()
		switch b {
		case "1":
			m.note -= int(p.strengh) / 2
			fmt.Printf("%s a attaqué %s de %v point de degats \nIl est maintenant à %v/%v\n", p.name, m.name, int(p.strengh)/2, m.note, m.notemax)
		case "2": // A revoir
			if p.energy-bos["python"] >= 0 {
				p.energy -= bos["python"]
				m.note -= int(p.strengh)
				fmt.Printf("%s a attaqué %s de %v point de degats \nIl est maintenant à %v/%v\n", p.name, m.name, p.strengh, m.note, m.notemax)
			} else {
				fmt.Println("Tu n'a plus assez d'energie")
				Enter()
				m.Training(p)
			}

		case "3":
			if p.IsInSkill("go") {
				if rand.Float64() < 5.0/p.agilite { //1 chance sur deux augmenter s'il ameliore une crtn stat story telling animated
					if p.energy-bos["go"] >= 0 {
						p.energy -= bos["go"]
						m.note -= int(p.strengh) * 2
						fmt.Printf("%s a attaqué %s de %v point de degats \nIl est maintenant à %v/%v\n", p.name, m.name, p.strengh*2, m.note, m.notemax)
					}
				} else {
					fmt.Println("Tu as raté ton coup")
				}
			} else {
				fmt.Println("Veuillez choisir une option valide")
				Enter()
				m.Training(p)
			}
		default:
			fmt.Println("Veuillez choisir une option valide")
			Enter()
			m.Training(p)
		}

	case "2":
		p.AccessInventory(2) //2=J'utilise un item en plein combat
	default:
		fmt.Println("Veuillez choisir une option valide")
		Enter()
		m.Training(p)
	}
}

func (m *Mentor) Training(p *Personnage) {

	var count int = 1
	fmt.Println("Vous êtes rentrés dans le tournoi \nVous êtes au ", round, "round")
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
		fmt.Println("Votre échauffement est maintenant terminé, vous avez gagné")
		p.exp += m.exp
		p.initiative += m.initiative
		p.wallet += m.wallet
		p.RandomObjects(0.5)
		fmt.Println("Vous avez battue ", m.name, " vous gagnez +", m.exp, " exp\n+", m.initiative, " d'initiative\n+", m.wallet, "€")
		if p.LimitSpace() {
			p.AddInventory("graine")
		}
		p.LevelUp()
		round--
		p.Display()
	} else {
		fmt.Println("Votre échauffement est maintenant terminé, vous avez perdu")
		round = 4
	}
	Enter()
	p.Menu()
}
