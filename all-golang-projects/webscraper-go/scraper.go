package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// this program scrapes data from a shopping site
// it will get name, price and the image url!

type Item struct {
	Name   string `json:"name"` // (not a very accurate description) define json fields (what data will look like in json file)
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	// c := colly.NewCollector() - instace for webscraper to work using colly package
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"), // pretty self explanitory, only allowed to scrape this domain
	)

	var items []Item // blank slice

	// onhtml = whenever we find this element (the param in quotation marks) in html code, run this:
	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		storeitem := Item{
			Name: h.ChildText("h2.product-title"),
			Price: h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"), // get source of every image
		}

		// use ChildText or it will print all the text in the website
		// (but replaces with spaces) and the things we need not blanked out
		
		// append data to a slice
		items = append(items, storeitem)

		//// fmt.Println(storeitem)
		//// h.ChildText("h2.product-title") // print specific part of text in  and only header2's
	})

	// website has multiple pages, so here we go again!
	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href")) // the actual code for the website doesnt have a full link inside this gets the full link in the href
		c.Visit(next_page) // loops through every page
	})

	// shows url as webscraper goes through it i think
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})


	c.Visit("https://j2store.net/demo/index.php/shop") // provide whole url for webscraper to visit

	content, err := json.Marshal(items)

	if err != nil{
		fmt.Println(err.Error())
	}

	// stores data in a file, 0644 is some thing that u dont need to understand
	os.WriteFile("products.json", content, 0644)

	//// fmt.Println(items) // print all data
}

