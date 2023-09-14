package main

import (
	pc "piscine/caracteristique"
)

func main() {
	var p1 pc.Personnage
	p1.Init("test", "test")
	p1.Display()
	p1.AccessInventory()
	p1.TakePot()
}
