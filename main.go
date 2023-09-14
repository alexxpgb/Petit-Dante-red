package main

type Personnage struct {
	name       string
	classe     string
	niveau     int
	PDVmax     int
	PDV        int
	inventaire []string
}

func main() {
	p1 = Personnage()
	p1.create("Alex", "Chiant")
}
