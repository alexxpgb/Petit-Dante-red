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
	notemax    float64 //infra
	note       int
	inventaire map[string]int
	leninv     int //taille inventaire
	skills     []string
	wallet     int
	armure     Equipment
	initiative float64 //dev
	strengh    float64 //cyber
	exp        int
	expmax     float64
	energy     int //mana
	intmax     float64
	agilite    float64 //IA Data
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
	TermPrint("❖ Quel est ton nom ?", 0, 0, termbox.ColorDefault)
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
					fmt.Print("\033[H\033[2J")
					fmt.Println("Tu veux supprimer quoi là?")
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
		fmt.Println("Pseudo non accepté")
		p.name = ""
		Enter()
		ReadInputO()
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
	p.agilite = 10
	p.armure.head = "rien"
	p.armure.body = "rien"
	p.armure.hand = "rien"
	p.agilite = 5
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
				if len(s) > 0 {
					x -= runewidth.RuneWidth(p.LastRune(s))
					s = s[:len(s)-1]
					tbprint(x, y, color, termbox.ColorDefault, ' ')
				} else {
					fmt.Print("\033[H\033[2J")
					fmt.Println("Tu veux supprimer quoi là?")
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
		p.expmax *= 1.5
		if p.niveau == "B1" {
			p.niveau = "B2"
			p.notemax *= 1.3
			p.initiative *= 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
		} else if p.niveau == "B2" {
			p.niveau = "B3"
			p.class()
		} else if p.niveau == "B3" {
			p.niveau = "M1"
			p.notemax *= 1.3
			p.initiative *= 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
		} else if p.niveau == "M1" {
			p.niveau = "M2"
			p.notemax *= 1.3
			p.initiative *= 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
			fmt.Println("Vous etes au niveau max")
			Enter()
		}
		fmt.Println("Bravo vous avez level up!!!!\nVous êtes maintenant en ", p.niveau)
		fmt.Printf("\nVous avez évolué voici votre niveau actuel\nNote max : %v←%v\nInitiative : : %v←%v\nForce : %v←%v\nForce vital : %v←%v\nAgilité : %v←%v\n\n\n", p.notemax/1.3, p.notemax, p.initiative/1.3, p.initiative, p.strengh/2, p.strengh, p.intmax/1.3, p.intmax, p.agilite/1.3, p.agilite)
		p.Display()
	}
}

func (p *Personnage) class() {
	if p.niveau == "B3" {
		fmt.Println("❖ Quelle est ta specialisation ?")
		fmt.Println("-------------------------------")
		fmt.Println("1/IA data (+ agilite) \n2/infra(+ notemax) \n3/cybersécurité(+ strengh) \n4/dev (+ initiative) ")
		fmt.Print("-------------------------------\n\n")
		fmt.Print("❖ Saisie le numéro de ta spécialité\n\n\n\n\n")
		answer := Scan()
		if answer == "1" { //Mettre des niveau different selon la classe
			p.classe = "IA data"
			p.notemax *= 1.3
			p.initiative *= 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 4
		} else if answer == "2" {
			p.classe = "infra "
			p.notemax *= 4
			p.initiative *= 1.3
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
		} else if answer == "3" {
			p.classe = "cybersécurité"
			p.notemax *= 1.3
			p.initiative *= 1.3
			p.strengh *= 4
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
		} else if answer == "4" {
			p.classe = "dev"
			p.notemax *= 1.3
			p.initiative *= 4
			p.strengh *= 2
			p.expmax *= 1.5
			p.intmax *= 1.3
			p.agilite *= 1.3
		}
		p.Display()
	}
}
