package main

import (
	"fmt"
	"time"
)

//START OMIT
func grita(baboseira string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(baboseira)
	}
}

func main() {
	go grita("firework")

	grita("conq")
}
//END OMIT