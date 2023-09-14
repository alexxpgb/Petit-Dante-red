package piscine

func (p Personnage) init(name string, classe string) {
	p.name := name
	p.classe := classe
	p.niveau := 1
	p.PDVmax := 10
	p.PDV := 10
	p.inventaire := []string
}
