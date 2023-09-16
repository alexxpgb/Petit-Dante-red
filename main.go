package main

import (
	pc "piscine/caracteristique"
)

func main() {
	var p1 pc.Personnage
	pc.Graphisme()
	p1.Init("Nath", "Info")
	p1.Menu()
}
