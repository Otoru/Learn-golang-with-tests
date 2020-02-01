package main

import "fmt"

const englishGreeting = "Hello "
const spanishGreeting = "Hola "
const frenchGreeting = "Bonjour "

// English const used to call Hello function in english
const English = "English"

// Spanish const used to call Hello function in spanish
const Spanish = "Spanish"

// French const used to call Hello function in french
const French = "French"

const world = "World"

// Hello function returns a greeting to the informed people
func Hello(people, language string) string {

	var greeting, name string

	switch people {
	case "":
		name = world
	default:
		name = people
	}

	switch language {
	case Spanish:
		greeting = spanishGreeting
	case French:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello("Vitor", "English"))
}
