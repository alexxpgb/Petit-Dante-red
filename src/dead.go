package piscine

import (
	"fmt"
	"time"
)

func (p Personnage) IsAlive() bool { //pas compliqué a comprendre celui la
	return p.note > 0 //true si p.note>0 false sinon
}

func (p *Personnage) Redouble() {
	if p.note <= 0 && p.IsInInventory("totem") { //Y a une fonction ou si on meurt on résucite une fois donc j'ai crée un objet totem et que si tu l'a tu peux résuciter
		p.note = 30
		fmt.Println("tu as une seconde chance")
		p.RemoveInventory("totem") //Une fois utilisé il disparait
	}
}

func (m *Mentor) Poison() { //Mentor en parametre car s'est lui qui prend le poison et pas moi
	for i := 0; i < 3; i++ {
		m.note -= 10
		fmt.Printf("๑Aie %s est à %d/%d\n", m.name, m.note, m.notemax)
		time.Sleep(time.Second * 1) //Toute les secondes sa fait un coup
	}
}

func (p *Personnage) TakePot(nb int) {
	var a bool
	for cle := range p.inventaire {
		if cle == "sucette" { //On parcours l'inventaire et si on a la sucette on l'utilise
			if p.note > int(p.notemax)-20 { //La sucette rapporte un +20 à ta note donc si tu l'utilise et que tu deborde sur ta note max y a un affichage diff
				fmt.Println("❖ Es-tu sur de vouloir utiliser la sucette")
				fmt.Println("1 pour oui")
				fmt.Print("2 pour non\n\n\n\n")
				switch Scan() {
				case "1":
					p.inventaire["sucette"] -= 1      //Je lui enlève une sucette dans mon inventaire
					if p.inventaire["sucette"] == 0 { //Si j'en ai plus je la supprime de mon inventaire
						delete(p.inventaire, "sucette")
					}
					p.note = int(p.notemax) //Dans le cas où ma note est superieur à ma note max - ma sucette donc ma note est au max car ma sucette fait un +20
					fmt.Println("Tu as pris une sucette tu es maintenant à", p.note)
					a = true
				case "2":
					fmt.Println("Fais plus attention la prochaine fois")
					a = true
				default:
					fmt.Println(" Peux tu répéter ?")
					p.TakePot(nb) //On le relance

				}
			} else {
				p.inventaire["sucette"] -= 1
				if p.inventaire["sucette"] == 0 {
					delete(p.inventaire, "sucette")
				}
				p.note += 20
				fmt.Println("Tu as pris une sucette tu es maintenant à", p.note)
				a = true
			}
		}
	}
	if !a { //Le cas contraire de a donc si on utilise une sucette on peut pas rentrer dans cette boucle
		fmt.Println("Tu n'as pas de sucette") // j'ai regardé partout dans mon inventaire mais tu n'as pas de sucette
	}
	if nb == 2 {
		return
	}
}
