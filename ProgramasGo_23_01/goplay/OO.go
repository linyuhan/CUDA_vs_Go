package main

import (
    "fmt"
)

type MinhaEstrutura struct {
    nome   string
    número int
}

func (minha *MinhaEstrutura) ImprimaNomeENúmero() {
    fmt.Printf("Nome:\t%s\nNúmero:\t%d\n", minha.nome, minha.número)
}

func main() {
    m := new(MinhaEstrutura)
    m.nome = "Fulano"
    m.número = 10
    m.ImprimaNomeENúmero()

    /* saída:
       Nome:   Fulano
       Número: 10
    */
}