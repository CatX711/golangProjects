/*
package main

import (
	"fmt"
	"time"
)

func main() {

	var userCache, cookies, Ox11 int
	var data = []*int{&userCache, &cookies, &Ox11} // array containing pointers to vars

	fmt.Println(data)
	tempdata(data, 4)
	fmt.Println(data)
}

func tempdata(vars []*int, wait int) {

	go func() {
		for _, v := range vars {
			time.Sleep(time.Duration(wait) * time.Second) // waits specified time
			*v = 0                                        // sets vars to null
		}
	}()
}
*/

package main

import (
	"fmt"
	"go-tempdata/temp"
	"go-tempdata/temp/mod" // auto imported
	"time"
)

func main() {
	var userCache, cookies, Ox11 int
	userCache = 86
	cookies = 339183019
	var data = []*int{&userCache, &cookies, &Ox11} // array containing pointers to vars

	name := "daniel" 
	num := 36

	fmt.Println(name)
	fmt.Println("Before:", userCache, cookies, Ox11)

	fmt.Println("blablabla")
	time.Sleep(2 * time.Second)
	fmt.Println("knljdlajdiajwd")
	
	temp.Tempdata(data, 1)

	fmt.Println("After:", userCache, cookies, Ox11)

	fmt.Println("Before:", name)

	mod.ModString(&name, "blabla")

	fmt.Println("After:", name)

	fmt.Println("Before:", num)

	mod.ModInt(&num, 4)
	
	fmt.Println("After:", num)
}



