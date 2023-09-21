package piscine

import "fmt"

func (p *Personnage) AddInventory(s string) { // quand t'ajoute un item kékiçepasse ?
	for cle := range p.inventaire {
		if cle == s {
			p.inventaire[cle]++ // si je l'ai deja on augmente de 1 le stock
			return
		}
	}
	p.inventaire[s] = 1 // si je l'ai pas je vais le chercher chez ta mère et je l'initialise
}

func (p *Personnage) RemoveInventory(s string) { // la c'est quand on enleve de l'inventaire
	for cle := range p.inventaire {
		if cle == s {
			delete(p.inventaire, cle) // ca supprime direct ! comme ton ex avec ton num
		}
	}
}

func (p *Personnage) IsInInventory(s string) bool { // on regarde si c'est dans l'inventaire ou pas
	for cle := range p.inventaire {
		if cle == s {
			return true // si ca y'est tu peux te le mettre dans le trou
		}
	}
	return false
}

func (p *Personnage) IsInSkill(s string) bool { // la on regarde si c'est dans ta skill list
	for _, c := range p.skills {
		if c == s {
			return true // pareil pour ton trou
		}
	}
	return false
}

func (p Personnage) LimitSpace() bool { // vu que tu fais pas de sport tu peux pas trop porter de truc donc on va bien regarder et choisir ce que tu peux porter sale faible
	var count int
	for cle := range p.inventaire {
		if cle != "" {
			count++
		}
	}
	if count > 10 {
		fmt.Println("Vous avez déjà 10 objets dans votre inventaire") //trop faible...
		return false
	}
	return true
}
func IsUpper(s string) bool { // on verifies si la lettre est majuscule (t'aurais pu utiliser Capitalise PD)
	nbs := len(s)
	nb := 0
	for _, c := range s {
		for i := 65; i < 91; i++ {
			if c == rune(i) {
				nb += 1
			}
			if nb == nbs {
				return true
			}
		}
	}
	return false
}

func IsLower(s string) bool { //(Pk pas Capitalise sale nègre ?)
	nbs := len(s)
	nb := 0
	for _, c := range s {
		for i := 97; i < 123; i++ {
			if c == rune(i) {
				nb += 1
			}
			if nb == nbs {
				return true
			}
		}
	}
	return false
}

func IsAlpha(s string) bool { // male Alpha mes couilles
	for _, c := range s {
		if c < 48 {
			return false
		}
		if c > 122 {
			return false
		}
	}
	return true
}

func (p *Personnage) UseObject(s string) { // si on utilise pas ca reste dans le stock et si ca reste... ca périme et après.. C'EST DEGEULASSE !
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
func ToLower(s string) string { // Capitalize ?
	var listf string
	for _, c := range s {
		if c > 64 && c < 91 {
			listf = listf + string(c+32)
		} else {
			listf = listf + string(c)
		}
	}
	return listf
}

func ToUpper(s string) string { // Je vais me répetter mais Capitalize?
	var listf string
	for _, c := range s {
		if c > 96 && c < 123 {
			listf = listf + string(c-32)
		} else {
			listf = listf + string(c)
		}
	}
	return listf
}

// bande de merde, vous pouvez utiliser capitalize !!!
