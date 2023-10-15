package main

import (
	"fmt"
)

func main() {

	s := "hElloworld"
	counter := 0

	for _, character := range s {

		switch character {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			counter++
		}

		//character := strings.ToLower(string(character))
		//if character == "a" || character == "e" || character == "i" || character == "o" || character == "u" {
		//	counter++
		//}
	}

	fmt.Println(counter)

}
