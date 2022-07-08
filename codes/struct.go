package main

import "fmt"

//START OMIT
type Crianca interface {
	JogaBolaDePapel()
}

type Hideo struct {
	bolinhas int
}

func (h Hideo) JogaBolaDePapel() {
	message := fmt.Sprintf("Hideo tinha %d bolinhas, agora tem %d", h.bolinhas, h.bolinhas - 1)
	fmt.Println(message);
}

func main() {
	hideo := Hideo{2}

	hideo.JogaBolaDePapel()
}
//END OMIT