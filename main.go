package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	Greeting()
	advanceGreeting("Sidd")
}

func Greeting() {
	fmt.Println(("Greeting Function"))
}

func advanceGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
