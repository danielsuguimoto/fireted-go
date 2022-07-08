package main

import "fmt"

//START OMIT
func saySomething(something string) {
	fmt.Println(something);
}

func getHelloMessage(name string) string {
	return fmt.Sprintf("Hello, %s!")
}

func main() {
	message := getHelloMessage("Hideo cabeludo")
	saySomething(message)
}
//END OMIT