package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoluxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bwm7 struct{}

func (b bwm7) ligaturbo() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoluxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
