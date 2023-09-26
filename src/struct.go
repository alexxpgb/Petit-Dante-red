package piscine

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	term "github.com/nsf/termbox-go"
)

type Personnage struct {
	name       string
	classe     string
	niveau     string
	notemax    float64
	note       int
	inventaire map[string]int
	leninv     int //taille inventaire
	skills     []string
	wallet     int
	armure     Equipment
	initiative float64
	strengh    float64
	exp        int
	expmax     float64
	energy     int //mana
	intmax     float64
}

type Equipment struct {
	head string
	body string
	hand string
}

type Mentor struct {
	name       string
	notemax    int
	note       int
	strengh    int
	wallet     int
	initiative float64
	exp        int
}

func (m *Mentor) InitMentor() {
	m.name = "Mentor"
	m.note = 100
	m.notemax = 100
	m.strengh = 10
	m.initiative = 20
	m.wallet = 100
}

func tbprint(x, y int, fg, bg termbox.Attribute, c rune) {
	termbox.SetCell(x, y, c, fg, bg)
	x += runewidth.RuneWidth(c)
	termbox.Flush()
}

func (p *Personnage) Init() { //Pour demander et luo attribuer le nom du personnage et ses infos

	x := 25
	y := 0
	var graph [][]rune

	graph = append(graph, []rune("❖ Quel est ton nom ?"))
	for i := range graph {
		x := 0
		for _, char := range graph[i] {
			tbprint(x, y+i, termbox.ColorDefault, termbox.ColorDefault, char)
			x += runewidth.RuneWidth(char)
		}
	}
	inputchek := false
	for inputchek == false {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEnter:
				inputchek = true
			case term.KeyBackspace:
				if len(p.name) > 0 {
					x -= runewidth.RuneWidth(p.LastRune(p.name))
					p.name = p.name[:len(p.name)-1]
					tbprint(x, y, termbox.ColorCyan, termbox.ColorDefault, ' ')
				} else {
					fmt.Println("Tu veut supprimer quoi la?")
				}

			default:
				if IsAlpha(string(ev.Ch)) {
					x = p.inputs(ev.Ch, x, y)
					p.name += string(ev.Ch)
				}
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
	if IsNum(p.name) {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Pseudo non acceptée")
		p.name = ""
		p.Init()
	}
	p.name = Capitalize(p.name)
	p.classe = "info"
	p.niveau = "B1"
	p.notemax = 100
	p.note = 50
	p.inventaire = map[string]int{"sucette": 2, "totem": 1}
	p.leninv = 10
	p.skills = []string{"python"}
	p.wallet = 50
	p.strengh = 5
	p.energy = 50
	p.intmax = 100
	p.initiative = 10
	p.expmax = 65
	p.Menu()

}

func Scan() string { //comme fmt.Scan sauf que ça te l'affiche aussi
	var graph [][]rune

	graph = append(graph, []rune("╰┈➤"))
	y := 40

	for i := range graph {
		x := 6
		for _, char := range graph[i] {
			tbprint(x, y+i, termbox.ColorDefault, termbox.ColorDefault, char)
			x += runewidth.RuneWidth(char)
		}
	}
	x := 11
	y = 40
	var s string
	inputchek := false
	for inputchek == false {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeySpace:
				x = p.inputs(' ', x, y)
				s += string(' ')
			case term.KeyEnter:
				inputchek = true
			case term.KeyBackspace:
				if len(p.name) > 0 {
					x -= runewidth.RuneWidth(p.LastRune(s))
					s = s[:len(s)-1]
					tbprint(x, y, color, termbox.ColorDefault, ' ')
				} else {
					fmt.Println("Tu veut supprimer quoi la?")
				}

			default:
				if IsAlpha(string(ev.Ch)) {
					x = p.inputs(ev.Ch, x, y)
					s += string(ev.Ch)
				}
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
	term.Clear(color, termbox.ColorDefault)
	return s
}

func (p *Personnage) inputs(input rune, x, y int) int {
	tbprint(x, y, color, termbox.ColorDefault, input)
	x += runewidth.RuneWidth(input)
	return x
}

func (p *Personnage) LevelUp() {
	if p.exp >= int(p.expmax) {
		p.exp = int(p.expmax) - p.exp
		p.expmax *= 1/2 + 1
		if p.niveau == "B1" {
			p.niveau = "B2"
			p.notemax *= 1.3
			p.initiative += 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
		} else if p.niveau == "B2" {
			p.niveau = "B3"
			p.class()
		} else if p.niveau == "B3" {
			p.niveau = "M1"
			p.notemax += 60
			p.initiative += 60
			p.strengh *= 4
			p.expmax += 100
			p.intmax += 60
		} else if p.niveau == "M1" {
			p.niveau = "M2"
			p.notemax += 120
			p.initiative += 120
			p.strengh *= 8
			p.expmax += 200
			p.intmax += 120
			fmt.Println("Vous etes au niveau max")
		}
	}
}

// d
// o
// n
// t
// f
// o
// r
// g
// e
// t
func (p *Personnage) class() {
	if p.niveau == "B3" {
		fmt.Println("❖ Quelle est ta specialisation ?")
		fmt.Println("-------------------------------")
		fmt.Println("1-IA data  2-infra 3-cybersécurité 4- dev  ")
		fmt.Println("❖ Saisie le numéro de ta spécialité")
		answer := Scan()
		if answer == "1" { //Mettre des niveau different selon la classe
			p.classe = "IA data"
		} else if answer == "2" {
			p.classe = "infra "
		} else if answer == "3" {
			p.classe = "cybersécurité"
		} else if answer == "4" {
			p.classe = "dev"
		}
	}
}
