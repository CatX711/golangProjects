/*package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("receipts.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// file.WriteString("testing")
	var month string
	var monthamount int
	var deleteChoice string
	var quit string
	var show string
	var total int


	fmt.Println("\n\nReset file? (yes/no) ")
	fmt.Scan(&deleteChoice)
	// file reset process
	if deleteChoice == "yes" {
		err := os.Remove("receipts.txt")
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	
		os.Create("receipts.txt") // Recreate the file
	}

	for {
		fmt.Println("\n\n\nEnter month: ")
		fmt.Scan(&month)
		fmt.Println("Enter amount: ")
		fmt.Scan(&monthamount)
		total += monthamount // adds to total
		pounds := strconv.Itoa(monthamount) // converts int to string
		file.WriteString(month + " " + pounds + "\n")

		fmt.Println("Quit program? (yes/no) ")
		fmt.Scan(&quit)

		if quit == "yes"{
			break
		} else if quit == "no"{
			continue // goes back to beginning of for loop
		} else{
			fmt.Println("ERROR: Unknown input.")
		}

	}

	fmt.Println("\n\n\n\nShow results? (yes/no) ")

	if show == "yes"{
		fmt.Println("Total: ", total)
	}	
}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("receipts.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var deleteChoice string
	var show string
	var total int

	fmt.Println("\n\nReset file? (yes/no) ")
	fmt.Scan(&deleteChoice)

	// file reset process
	if deleteChoice == "yes" {
		err := os.Remove("receipts.txt")
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}

		os.Create("receipts.txt") // Recreate the file
	}

	// Reading and calculating total from existing file contents
	receiptsFile, err := os.Open("receipts.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer receiptsFile.Close()

	scanner := bufio.NewScanner(receiptsFile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Split line into words
		if len(words) >= 2 {
			amount, err := strconv.Atoi(words[len(words)-1]) // Convert the last word to integer
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			total += amount // Add the extracted amount to total
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	var quit string

	for {
		var month string
		var monthAmount int

		fmt.Println("\n\n\nEnter month: ")
		fmt.Scan(&month)
		fmt.Println("Enter amount: ")
		fmt.Scan(&monthAmount)

		total += monthAmount // adds to total
		pounds := strconv.Itoa(monthAmount) // converts int to string
		file.WriteString(month + " " + pounds + "\n")

		fmt.Println("Quit program? (yes/no) ")
		fmt.Scan(&quit)

		if quit == "yes" {
			break
		} else if quit == "no" {
			continue // goes back to beginning of for loop
		} else {
			fmt.Println("ERROR: Unknown input.")
		}
	}

	// Print the total
	fmt.Println("\n\n\n\nShow results? (yes/no) ")
	fmt.Scan(&show)

	if show == "yes" {
		fmt.Println("Total: ", total)
	}
}
