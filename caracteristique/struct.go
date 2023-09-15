package piscine

type Personnage struct {
	name       string
	classe     string
	niveau     string
	notemax    int
	note       int
	inventaire map[string]int
}
