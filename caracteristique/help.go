package piscine

import "fmt"

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

func (p *Personnage) IsInSkill(s string) bool {
	for _, c := range p.skills {
		if c == s {
			return true
		}
	}
	return false
}

func (p Personnage) LimitSpace() bool {
	var count int
	for cle := range p.inventaire {
		if cle != "" {
			count++
		}
	}
	if count > 10 {
		fmt.Println("Vous avez déjà 10 objets dans votre inventaire")
		return false
	}
	return true
}

func (p *Personnage) UseObject(s string) {
	for cle := range p.inventaire {
		if cle == s {
			p.inventaire[cle]--
			if cle == "sucette" {
				p.TakePot()
			}
			if cle == "totem" {
				fmt.Println("Vous ne pouvez utiliser ce totem que si vous mourrez")
			}
			if cle == "Skill: go" {
				fmt.Println("Vous pouvez maintenant apprendre le skill go dans Book Of Skill")
				p.RemoveInventory("Skill: go")
				p.BookOfSkills("go")
			}
		} else {
			fmt.Println("Vous n'avez pas cette object dans votre inventaire")
		}
	}
}
