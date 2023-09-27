package piscine

import (
	"fmt"
	"math/rand"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	term "github.com/nsf/termbox-go"
)

func (p *Personnage) UseObject(m *Mentor, s string, nb int) { // si on utilise pas ca reste dans le stock et si ca reste... ca périme et après.. C'EST DEGEULASSE !
	for cle := range p.inventaire {
		if cle == s {
			if cle == "sucette" {
				p.TakePot(nb)
				return
			}
			if cle == "totem" {
				fmt.Println("Vous ne pouvez utiliser ce totem que si vous mourrez")
				return
			}
			if cle == "Skill: go" {
				fmt.Println("Vous pouvez maintenant apprendre le skill go dans Book Of Skill")
				p.RemoveInventory("Skill: go")
				AppendSkill("go", 7)
				return
			}
			if cle == "Upgrade inventaire" {
				p.inventaire[cle]--
				p.UpgradeInventory()
				fmt.Printf("Vous avez aggrandi votre inventaire maintenant vous avez jusqu'a %d places disponible\n", p.leninv)
				return
			}
			if cle == "douche" {
				if nb == 2 { //Le nb c'est pour savoir si ma commande vient du menu ou d'un fight
					m.Poison()
					p.RemoveInventory("douche")
				} else { //Dans ce cas la s'il vient du menu c'est qu'il essaye cette objet sauf qu'il y a personne donc ca fait rien dans l'autre cas c'est sur l'adversaire
					fmt.Println("La douche est dangereuse dans ce jeu ,rappelle toi que tu est en info")
				}
				return
			}
			if cle == "skittles" {
				p.TakeInt(nb)
				return
			}
		}
	}
	fmt.Println("Vous n'avez pas cette objet dans votre inventaire")
}

func (p *Personnage) AddInventory(s string) { // quand t'ajoute un item kékiçepasse ?
	for cle := range p.inventaire {
		if cle == s {
			p.inventaire[cle]++ // si je l'ai deja on augmente de 1 le stock
			return
		}
	}
	p.inventaire[s] = 1 // si je l'ai pas je vais le chercher chez ta mère et je l'initialise
}

func TransvalseList(tab map[string]int) []string { //cast une map en liste
	var lst []string
	for cle := range tab {
		lst = append(lst, cle)
	}
	return lst
}
func TermPrint(s string, x, y int, color termbox.Attribute) {
	var graph [][]rune
	graph = append(graph, []rune(s))
	for i := range graph {
		for _, char := range graph[i] {
			tbprint(x, y+i, color, termbox.ColorDefault, char)
			x += runewidth.RuneWidth(char)
		}
	}
}
func (p *Personnage) RemoveInventory(s string) { // la c'est quand on enleve de l'inventaire
	for cle := range p.inventaire {
		if cle == s {
			p.inventaire[cle]--
			if p.inventaire[cle] == 0 {
				delete(p.inventaire, cle) // ca supprime direct ! comme ton ex avec ton num
			}
		}
	}
}

func (p Personnage) IsInInventory(s string) bool { // on regarde si c'est dans l'inventaire ou pas
	for cle := range p.inventaire {
		if cle == s {
			return true // si ca y'est tu peux te le mettre dans le trou
		}
	}
	return false
}
func IsInMap(m map[string]int, s string) bool { // on regarde si c'est dans l'inventaire ou pas
	for cle := range m {
		if cle == s {
			return true // si ca y'est tu peux te le mettre dans le trou
		}
	}
	return false
}
func IsInList(s string, lst []string) bool {
	for _, l := range lst {
		if l == s {
			return true
		}
	}
	return false
}
func (p Personnage) IsInSkill(s string) bool { // la on regarde si c'est dans ta skill list
	for _, c := range p.skills {
		if c == s {
			return true // pareil pour ton trou
		}
	}
	return false
}

func Len(inv map[string]int) int {
	var count int
	for _, s := range inv {
		count++
		s++
	}
	return count
}

func (p Personnage) LimitSpace() bool { // vu que tu fais pas de sport tu peux pas trop porter de truc donc on va bien regarder et choisir ce que tu peux porter sale faible
	var count int
	for cle := range p.inventaire {
		if cle != "" {
			count++ //pour avoir un len
		}
	}
	if count > p.leninv { //J'ai crée un element dans la structure pour verifier si on a la place
		fmt.Println("Vous n'avez plus de place dans votre inventaire") //trop faible...
		return false
	}
	return true
}

func (p *Personnage) UpgradeInventory() {
	if p.leninv < 40 { //Car c'est le max qu'il veulent
		p.leninv += 10
	}
}

func IsUpper(s string) bool {
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

func IsLower(s string) bool {
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

func IsAlpha(s string) bool {
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

func ToLower(s string) string {
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

func IsNum(str string) bool {
	for _, chr := range str {
		if chr > 47 && chr < 58 {
			return true
		}
	}
	return false
}

func ToUpper(s string) string {
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

func (p *Personnage) LastRune(s string) rune { //Même moi je sais plus comment sa marche elle est trop compliqué ce package
	var count int = 0
	for _, c := range s {
		if count == len(s)-1 {
			return c
		}
	}
	return '0'
}

func Capitalize(s string) string {
	var new string
	if s == "" {
		return ""
	}
	if IsLower(string(s[0])) {
		new += ToUpper(string(s[0]))
	} else {
		new += string(s[0])
	}
	for i, c := range s {
		if i == 0 {
			continue
		}
		if !IsAlpha(string(new[len(new)-1])) {
			if i == len(s) {
				break
			}
			if IsLower(string(c)) {
				new += ToUpper(string(c))
			} else {
				new += string(c)
			}
		} else if IsUpper(string(c)) {
			new += ToLower(string(c))
		} else {
			new += string(c)
		}
	}
	return new
}

func (p *Personnage) RandomObjects(nb float64) {
	rand := rand.Float64() * nb
	if rand >= 1 {
		rand = 0.9
	}
	if rand <= 0.3 {
		p.AddInventory("Bouteille en plastique")
		fmt.Println("Bravo vous avez gagné une bouteille en plastique")
	} else if rand == 0.4 {
		p.AddInventory("carte graphique")
	} else if rand == 0.5 {
		p.AddInventory("febreze")
	} else if rand == 0.6 {
		p.AddInventory("clavier mecanique")
	} else if rand == 0.7 {
		p.AddInventory("souris")
	}
}

func Enter() {
	var chck bool
	for !chck {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEnter:
				chck = true
			}
		}
	}
}
