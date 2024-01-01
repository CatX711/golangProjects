package meanie

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

// Badword struct to store swear words
type Badword struct {
	SWord string // swear word
}

var nameofMeanie string

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wiktionary.org"),
	)

	// stores swear words
	var words []Badword

	// OnHTML is triggered when the specified HTML element is found
	c.OnHTML(".mw-category-group ul li", func(e *colly.HTMLElement) {
		// extract the text of the element and store it in the Badword struct
		word := Badword{
			SWord: e.Text,
		}
		// append the word to the slice of swear words
		words = append(words, word)
	})

	// OnRequest is triggered every time a new request is made
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String()) // Print the URL being visited
	})

	// visit the URL
	c.Visit("https://en.wiktionary.org/wiki/Category:English_swear_words")

	// wait for collector to finish
	c.Wait()

	// printf to not go on newln
	fmt.Printf("What is the MEANIE's name? ")
	fmt.Scan(&nameofMeanie)

	// Print all the swear words
	for _, word := range words {
		fmt.Println(nameofMeanie, "is a", word.SWord)
		time.Sleep(100 * time.Millisecond)
	}
}
