package piscine

import "fmt"

func (m *Mentor) MentorPattern(p *Personnage) {
	for p.note > 0 || m.note > 0 {
		for i := 0; i < 3; i++ {
			p.note -= m.strengh
			fmt.Printf("%s a attaqué %s de %d point \nIl est maintenant à%d/%d\n", m.name, p.name, m.strengh, p.note, p.notemax)
		}
		m.strengh *= 2
	}
}

func (p *Personnage) PersonnagePattern(m *Mentor) {
	fmt.Println("Menu:")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire\n\n\n\n\n\n")
	a := Scan()
	switch a {
	case "1":
		fmt.Println("Avec quels skills veut tu attaquer?")
		for i, c := range p.skills {
			fmt.Println(i+1, c)
		}
		b := Scan()
		switch b {
		case "1":
			p.strengh = 7
		case "2":
			p.strengh = 10
		}
	case "2":
		p.AccessInventory()
	default:
		fmt.Println("Veuillez choisir une option valide")
	}
}

func (m *Mentor) CharTurn(p *Personnage) {
	for p.note > 0 || m.note > 0 {
		p.PersonnagePattern(m)
		m.MentorPattern(p)
	}
}
