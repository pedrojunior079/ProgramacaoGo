package main

import "fmt"

func main(){
    i := 1
	var p *int = nil 
	p = &i // pega o endereco da variavel
	*p++ // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmético de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}