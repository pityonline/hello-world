package main

import "fmt"

// say prints something
func say(s string) string {
	return s
}

func main() {
	fmt.Println(say("Hello world!"))
}
