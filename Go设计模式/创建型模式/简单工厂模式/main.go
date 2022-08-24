package main

import "fmt"

func main() {
	dog := NewAPI("dog")
	fmt.Println(dog.Say("goDevNice"))
}
