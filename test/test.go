package main

import (
	"fmt"
	"strings"
)

// Definición de la función fuera de Test
func firstLetterUpper(entity string) string {
	return strings.ToUpper(entity[:1]) + entity[1:]
}

func firstLetterLower(entity string) string {
	return strings.ToLower(entity[:1]) + entity[1:]
}

func Test() {
	salida := firstLetterUpper("entity")
	salidaMayusculs := firstLetterLower("Entity")
	fmt.Println(salida)
	fmt.Println(salidaMayusculs)
}

func main() {
	Test()
}
