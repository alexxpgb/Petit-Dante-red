package piscine

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	term "github.com/nsf/termbox-go"
)

type Personnage struct {
	name         string
	classe       string
	niveau       string
	notemax      int
	note         int
	inventaire   map[string]int
	leninv       int //taille inventaire
	skills       []string
	wallet       int
	armure       Equipment
	initiative   int
	strengh      int
	exp          int
	expmax       int
	intelligence int //mana
	intmax       int
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
	initiative int
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

	x := 3
	y := 5
	fmt.Println("❖ Quel est ton nom ?")
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
					tbprint(x, y, termbox.ColorBlack, termbox.ColorDefault, ' ')
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
	p.intelligence = 100
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
			case term.KeyEnter:
				inputchek = true
			case term.KeyBackspace:
				if len(p.name) > 0 {
					x -= runewidth.RuneWidth(p.LastRune(s))
					s = s[:len(s)-1]
					tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, ' ')
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
	term.Clear(termbox.ColorDefault, termbox.ColorDefault)
	return s
}

func (p *Personnage) inputs(input rune, x, y int) int {
	tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, input)
	x += runewidth.RuneWidth(input)
	return x
}

func (p *Personnage) class() {
	if p.niveau == "B3" {
		fmt.Println("❖ Quelle est ta specialisation ?")
		fmt.Println("-------------------------------")
		fmt.Println("1-IA data  2-infra 3-cybersécurité 4- dev  ")
		fmt.Println("❖ Saisie le numéro de ta spécialité")
		answer := Scan()
		if answer == "1" {
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
