package piscine

import "fmt"

func (p *Personnage) Redouble() {
	if p.note <= 0 && p.IsInInventory("totem") {
		p.note = 10
		fmt.Println("tu a une seconde chance")
		p.RemoveInventory("totem")
	}
}
