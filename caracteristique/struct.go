package piscine

type Personnage struct {
	name       string
	classe     string
	niveau     string
	notemax    int
	note       int
	inventaire map[string]int
}

func (p *Personnage) Init(name string, classe string) {
	p.name = name
	p.classe = classe
	p.niveau = "B1"
	p.notemax = 100
	p.note = 50
	p.inventaire = map[string]int{"sucette": 3, "totem": 1}
}
