package main

import (
	"fmt"
	"time"
)

//START OMIT
func fazNumeroDois(chTerminou chan int) {
	time.Sleep(10 * time.Second)
	chTerminou <- 1
}
func leJornal(chTerminou chan int, chPaginas chan int) {
	paginas := 0
	for _ = range time.Tick(2 * time.Second) {
		select {
		case <-chTerminou:
			chPaginas <- paginas
			return
		default:
			paginas++
		}
	}
}
func main() {
	chTerminou, chPaginas := make(chan int, 1), make(chan int, 1)
	go fazNumeroDois(chTerminou)
	go leJornal(chTerminou, chPaginas)
	fmt.Println("Páginas lidas: ", <-chPaginas)
}
//END OMIT