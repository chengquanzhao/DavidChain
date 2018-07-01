package main

import (
	"DavidChain/basicPrototype1/BLC"
	"fmt"
)

func main() {

	block := BLC.CreateBlock("Genenis Block ", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
}
