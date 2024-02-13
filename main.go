package main

import (
	"fmt"
	"log"
	"os"

	"github.com/munanadi/cc-json-parser/stack"
)

const (
	OPEN_CURLY  string = "OPEN_CURLY"
	CLOSE_CURLY string = "CLOSE_CURLY"
)

func main() {

	matchingPairs := map[byte]byte{
		'}': '{',
	}

	// Will start by using a stack and checking for brackets and quotes
	fileBytes, err := os.ReadFile("tests/step1/valid.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(fileBytes) == 0 {
		fmt.Println("Invalid JSON")
		return
	}

	stack := stack.NewStack(len(fileBytes))

	for _, v := range fileBytes {
		// fmt.Print(string(v))
		tokenType, err := checkIfToken(v)
		if err != nil {
			log.Fatal(err)
		}

		// If its valid token, Check for in the stack
		if tokenType == OPEN_CURLY || tokenType == CLOSE_CURLY {

			topElement := stack.Top()

			// Stack is empty, then add to the stack
			if stack.Length() == -1 {
				stack.AddToStack(v)
			}

			if topElement == matchingPairs[v] && topElement != byte(0) {
				// brackets match, pop stack
				_, _ = stack.RemoveFromStack()

				if stack.Length() == -1 {
					// Empty the stack
					_, _ = stack.RemoveFromStack()
				}
			}
		}
	}

	// stack.PrintStack()

	// Stack should be empty here for valid JSON
	if stack.Length() == -1 {
		fmt.Println("valid JSON")
		return
	} else {
		fmt.Println("Invalid JSON")
		return
	}
}

// checkIfToken will take an byte and return if its a token or not
// it will return what token it is
// or thrown an error
func checkIfToken(element byte) (string, error) {
	switch element {
	case '{':
		return OPEN_CURLY, nil
	case '}':
		return CLOSE_CURLY, nil
	default:
		return "ILLEGAL_CHARACTER", fmt.Errorf("dont know what character %v", element)
	}
}
