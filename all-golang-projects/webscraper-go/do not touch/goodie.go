package goodie

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Goodword struct to store positive or neutral words
type Goodword struct {
	GWord string // positive or neutral word
}

var nameOfKindPerson string

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.goodgoodgood.co"),
	)

	// stores positive or neutral words
	var words []Goodword

	// OnHTML is triggered when the specified HTML element is found
	c.OnHTML("table tr td", func(e *colly.HTMLElement) {
		// extract the text of the element and store it in the Goodword struct
		word := Goodword{
			GWord: e.Text,
		}
		// append the word to the slice of positive or neutral words
		words = append(words, word)
	})

	// OnRequest is triggered every time a new request is made
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String()) // Print the URL being visited
	})

	// visit the URL
	c.Visit("https://www.goodgoodgood.co/articles/positive-words")

	// wait for collector to finish
	c.Wait()

	// ask for the kind person's name
	fmt.Printf("What is the GOODIE's name? ")
	fmt.Scan(&nameOfKindPerson)

	// Print all the positive or neutral words
	for _, word := range words {
		fmt.Println(nameOfKindPerson, "is", word.GWord)
	}
}
