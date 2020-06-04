package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello function returns a greeting
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
