package main

import (
	"fmt"
	"strings"
)

//START OMIT
func swap(a, b string) (string, string) {
	return b, a
}

func join(parts ...string) (message string) {
	message = strings.Join(parts, " e ")
	return
}

func main() {
	primeiraVoz, segundaVoz := "Maiara", "Maraisa"

	fmt.Println(join(primeiraVoz, segundaVoz))

	primeiraVoz, segundaVoz = swap(primeiraVoz, segundaVoz)

	fmt.Println(join(primeiraVoz, segundaVoz))
}
//END OMIT