package main

import (
	"fmt"
	"strings"
	"time"
)

type Location struct {
	Continent  string
	Country    string
	CityOrTown string
	StreetName string
}

func main() {
	User := Location{}
	var proceed string

	for {
		fmt.Println("You have selected \"Fluffy Cat 04\" for purchase.")
		fmt.Println("Please enter the following billing information to continue with your order:")
		fmt.Printf("\n\nContinent: ")
		fmt.Scan(&User.Continent)
		fmt.Printf("\n\nCountry or State: ")
		fmt.Scan(&User.Country)
		fmt.Printf("\n\nCity or Town: ")
		fmt.Scan(&User.CityOrTown)
		fmt.Printf("\n\nStreet Name: ")
		fmt.Scan(&User.StreetName)

		fmt.Printf("\x1b[2J") // clears screen
		fmt.Println("\n\n\nYou have entered the following information:\nContinent:", User.Continent, "\nCountry/State:", User.Country, "\nCity/Town:", User.CityOrTown, "\nStreetname:", User.StreetName)
		fmt.Println("\nIs this correct? y/n")
		fmt.Scan(&proceed)

		// convert answer to lowercase
		if strings.ToLower(proceed) == "yes" {
			fmt.Println("\nThank you for purchasing \"Fluffy Cat 04\"!")
			break
		} else if strings.ToLower(proceed) == "no" {
			fmt.Println("Restarting process...")
			time.Sleep(3 * time.Second)

			fmt.Printf("\x1b[2J") // clears screen
			continue              // goes back to beginning
		} else {
			fmt.Println("Please try again")
		}
	}
}
