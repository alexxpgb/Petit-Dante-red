package piscine

func (p *Personnage) AddInventory(s string) {
	for cle := range p.inventaire {
		if cle == s {
			p.inventaire[cle]++
			return
		}
	}
	p.inventaire[s] = 1
}

func (p *Personnage) RemoveInventory(s string) {
	for cle := range p.inventaire {
		if cle == s {
			delete(p.inventaire, cle)
		}
	}
}

func (p *Personnage) IsInInventory(s string) bool {
	for cle := range p.inventaire {
		if cle == s {
			return true
		}
	}
	return false
}
