package piscine

import (
	"fmt"
	"time"
)

func (p *Personnage) Redouble() {
	if p.note <= 0 && p.IsInInventory("totem") { //Y a une fonction ou si on meurt on résucite une fois donc j'ai crée un objet totem et que si tu l'a tu peux résuciter
		p.note = 10
		fmt.Println("tu a une seconde chance")
		p.RemoveInventory("totem") //Une fois utilisé il disparait
	}
}

func (p *Personnage) Poison() {
	for i := 0; i < 3; i++ {
		p.note -= 10
		fmt.Printf("Aie tu est a %d/%d\n", p.note, p.notemax)
		time.Sleep(time.Second * 1) //Toute les secondes sa fait un coup
	}
}
