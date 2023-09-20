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
	notemax    int
	note       int
	inventaire map[string]int
	leninv     int
	skills     []string
	wallet     int
	armure     Equipment
}

type Equipment struct {
	head string
	body string
	foot string
}

type Mentor struct {
	name    string
	notemax int
	note    int
	strengh int
}

func tbprint(x, y int, fg, bg termbox.Attribute, c rune) {
	termbox.SetCell(x, y, c, fg, bg)
	x += runewidth.RuneWidth(c)
	termbox.Flush()
}

func (p *Personnage) Init() {
	x := 2
	y := 2
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
					tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, ' ')
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
	p.inventaire = map[string]int{"sucette": 3, "totem": 1}
	p.leninv = 10
	p.skills = []string{"python"}
	p.wallet = 50
	p.Menu()

}

func Scan() string {
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println("₪₪                                                                                                                               ₪₪")
	fmt.Println("₪₪                                                                                                                               ₪₪")
	fmt.Println("₪₪                                                                                                                               ₪₪")
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	x := 2
	y := 16
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
