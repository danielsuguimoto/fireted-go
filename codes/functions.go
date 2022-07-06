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
	message := getHelloMessage("Daniel")
	saySomething(message)
}
//END OMIT