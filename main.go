package main

import (
	"fmt"
	"log"
	"os"

	"github.com/munanadi/cc-json-parser/stack"
)

const (
	OPEN_CURLY    string = "OPEN_CURLY"
	CLOSE_CURLY   string = "CLOSE_CURLY"
	DOUBLE_QUOTES string = "DOUBLE_QUOTES"
	ALPHABET      string = "ALPHABET"
	NUMBER        string = "NUMBER"
	COLON         string = "COLON"
	COMMA         string = "COMMA"
)

func main() {

	// Will start by using a stack and checking for brackets and quotes
	fileBytes, err := os.ReadFile("tests/step2/valid2.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(fileBytes) == 0 {
		fmt.Println("Invalid JSON")
		return
	}

	fmt.Println(string(fileBytes))

	parsedStack := createStackFromBytes(fileBytes)

	parsedStack.PrintStack()

	/// I think Iam going to parse the input into tokens and put it into a []tokens array

	// for _, v := range fileBytes {

	// 	if v == '\n' || v == ' ' {
	// 		// Ignore new lines and empty spaces?
	// 		continue
	// 	}

	// 	tokenType, err := checkIfToken(v)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// If its valid token, Check for in the stack
	// 	if tokenType == OPEN_CURLY || tokenType == CLOSE_CURLY || tokenType == DOUBLE_QUOTES || tokenType == COLON || tokenType == COMMA {

	// 		// If : push into stack
	// 		if tokenType == COLON || tokenType == COMMA {
	// 			stack.AddToStack(v)
	// 		}

	// 		topElement := stack.Top()

	// 		if topElement == ',' && tokenType != DOUBLE_QUOTES {
	// 			stack.PrintStack()
	// 			fmt.Println(", not followed by \"")
	// 			break
	// 		}

	// 		if topElement == ':' && tokenType != DOUBLE_QUOTES {
	// 			stack.PrintStack()
	// 			fmt.Println(": not followed by \"")
	// 			break
	// 		}

	// 		// Stack is empty, then add to the stack
	// 		if stack.Length() == -1 {
	// 			stack.AddToStack(v)
	// 		}

	// 		if topElement == matchingPairs[v] && topElement != byte(0) {
	// 			// brackets match, pop stack
	// 			_, _ = stack.RemoveFromStack()

	// 			if stack.Length() == -1 {
	// 				// Empty the stack
	// 				_, _ = stack.RemoveFromStack()
	// 			}
	// 		}
	// 	} else if tokenType == ALPHABET || tokenType == NUMBER {
	// 		continue
	// 	}
	// }

	// // stack.PrintStack()

	// // Stack should be empty here for valid JSON
	// if stack.Length() == -1 {
	// 	fmt.Println("valid JSON")
	// 	return
	// } else {
	// 	fmt.Println("Invalid JSON")
	// 	return
	// }
}

func createStackFromBytes(fileBytes []byte) *stack.Stack[string] {
	stack := stack.NewStack[string](len(fileBytes))

	for _, v := range fileBytes {
		if v == '\n' || v == ' ' {
			continue
		}

		tokenType, err := checkIfToken(v)
		if err != nil {
			log.Fatal(err)
		}

		topElement := stack.Top()

		if topElement == ALPHABET && (tokenType == ALPHABET || tokenType == NUMBER) {
			// Alphanumeric together is okay
			continue
		} else {
			stack.AddToStack(tokenType)
		}
	}

	return stack
}

// checkIfToken will take an byte and return if its a token or not
// it will return what token it is
// or thrown an error
func checkIfToken(element byte) (string, error) {

	if element == '{' {
		return OPEN_CURLY, nil
	} else if element == '}' {
		return CLOSE_CURLY, nil
	} else if element == '"' {
		return DOUBLE_QUOTES, nil
	} else if (element > 64 && element < 91) || (element > 96 && element < 123) {
		return ALPHABET, nil
	} else if element > 47 && element < 58 {
		return NUMBER, nil
	} else if element == ':' {
		return COLON, nil
	} else if element == ',' {
		return COMMA, nil
	} else {
		return "ILLEGAL_CHARACTER", fmt.Errorf("dont know what character %v %s", element, string(element))
	}

}
