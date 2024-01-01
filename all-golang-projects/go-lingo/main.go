// used to work but it doesnt now whatever bro idc


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Custom function to interpret commands
func interpretCommand(input string, variables map[string]string) {
	// Check if the input starts with "--Req("
	if strings.HasPrefix(input, "--Req(") && strings.HasSuffix(input, ")") {
		content := input[len("--Req(") : len(input)-1] // Extracting content within parentheses
		if val, ok := variables[content]; ok {
			fmt.Println(val) // Print variable value if it exists
		} else {
			fmt.Println(content) // Print string if it's not a variable
		}
	} else {
		fmt.Println("Invalid command:", input)
	}
}

func main() {
	fileName := "example.li" // Change this to your file's name or provide the full path
	if strings.HasSuffix(fileName, ".li") {
		// Attempt to open the file
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		// Read file line by line
		scanner := bufio.NewScanner(f)
		variables := make(map[string]string)

		for scanner.Scan() {
			line := scanner.Text()

			// Ignore comments and empty lines
			if strings.HasPrefix(line, "@") || strings.TrimSpace(line) == "" {
				continue
			}

			// Parse the line to identify functions and variables
			// For simplicity, let's assume variables are declared as "varName = value"

			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				varName := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				variables[varName] = value
			} else {
				// Interpret command if it's not a variable declaration
				interpretCommand(line, variables)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}


/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Knetic/govaluate"
)

// Custom function to interpret commands
func interpretCommand(input string, variables map[string]string) {
	// Check if the input starts with "--Req("
	if strings.HasPrefix(input, "--Req(") && strings.HasSuffix(input, ")") {
		content := input[len("--Req(") : len(input)-1] // Extracting content within parentheses
		if val, ok := variables[content]; ok {
			fmt.Println(val) // Print variable value if it exists
		} else if isMathExpression(content) {
			result, err := evaluateMathExpression(content)
			if err != nil {
				fmt.Println("Error evaluating expression:", err)
			} else {
				fmt.Println(result) // Print the result of the mathematical expression
			}
		} else {
			fmt.Println(content) // Print string if it's not a variable or expression
		}
	} else {
		fmt.Println("Invalid command:", input)
	}
}

// Check if the input is a mathematical expression
func isMathExpression(input string) bool {
	expression, err := govaluate.NewEvaluableExpression(input)
	return err == nil && expression != nil
}

// Evaluate the mathematical expression
func evaluateMathExpression(expressionStr string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(expressionStr)
	if err != nil {
		return nil, err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	fileName := "example.li" // Change this to your file's name or provide the full path
	if strings.HasSuffix(fileName, ".li") {
		// Attempt to open the file
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		// Read file line by line
		scanner := bufio.NewScanner(f)
		variables := make(map[string]string)

		for scanner.Scan() {
			line := scanner.Text()

			// Ignore comments and empty lines
			if strings.HasPrefix(line, "@") || strings.TrimSpace(line) == "" {
				continue
			}

			// Parse the line to identify functions and variables
			// For simplicity, let's assume variables are declared as "varName = value"

			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				varName := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				variables[varName] = value
			} else {
				// Interpret command if it's not a variable declaration
				interpretCommand(line, variables)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}
*/