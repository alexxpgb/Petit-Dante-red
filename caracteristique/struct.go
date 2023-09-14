package piscine

type Personnage struct {
	name       string
	classe     string
	niveau     int
	PDVmax     int
	PDV        int
	inventaire map[string]int
}
