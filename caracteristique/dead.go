package piscine

import (
	"fmt"
	"time"
)

func (p *Personnage) Redouble() {
	if p.note <= 0 && p.IsInInventory("totem") {
		p.note = 10
		fmt.Println("tu a une seconde chance")
		p.RemoveInventory("totem")
	}
}

func (p *Personnage) Poison() {
	for i := 0; i < 3; i++ {
		p.note -= 10
		fmt.Printf("Aie tu est a %d/%d\n", p.note, p.notemax)
		time.Sleep(time.Second * 1)
	}
}
