package main

import "fmt"

// functions can be defined inside or outside the main function
func calc(x, y int) int { // shorten 	x int, y int	 to 	x, y 	int if they have a shared type
	return x + y
}

// structs:
type person struct { // must be outside main function
	name string
	age  int
}

/* this must be in main func but its here for demo purposes

func main(){
	p := person{name: "Daniel", age: 12}
	n := person{name: "James", age: 4}

	fmt.Println(p.name, p.age)
	fmt.Println(n.name, n.age)
}

output:

Daniel 12
James 4

! Structs are kinda like classes, they provide a structure for variables, if that makes sense lol
! Hopefully the example above is good enough for my brain to understand
*/

func main() {
	fmt.Printf("\n\n\n")

	var num1 int = 22
	num2 := 5 // type inferred
	sum := 0

	// add to the value of a variable:
	x := 1
	x++    // increases value by 1
	x += 5 // adds a specific value

	// if
	y := 24052
	if x < y {
		fmt.Println("True!")
	}
	if x == y { // always use == in an if statement
		fmt.Println("False!")
	}

	// use else like this:
	/*
		if num != 88 {
			fmt.Println("Num is not equal to 88!")
		} else{
			fmt.Println("Num is equal to 88!")
		}

		Or:

		if num != 88 {
			fmt.Println("Num is now not equal to 88!")
		} else if num == 89{
			fmt.Println("Num is now equal to 89!")
		} else{
			fmt.Println("Num is now equal to 88!")
		}

	*/

	// array
	var cars = [4]string{"Volvo", "BMW", "Ford", "Madza"} // length is defined as 4
	var cars2 = [...]string{"Tesla"}                      // length is inferred
	apples := [3]int{1, 2, 86}                            // type cannot be inferred even with type inference operator (:=)

	cars[0] = "Opel" // Change value of Volvo
	// (first element of array is 0, second is 1, third is 2, etcetera)

	// you can also define arrays without giving a value first
	var arr1 [3]string  // value is not defined
	// idk about this one: arr2 := [3]string{} // when making a blank slice or array using := you need {} following the array type

	arr1[0] = "Cool!" // first element of array is now "Cool!"
	// arr1 is now:
	// ["Cool!", "", ""]
	//    0      1   2     <--- index value

	// slice:
	var slice = []int{1, 2, 3, 4, 5, 6} // slices are important because with normal and inferred length arrays you cannot add more values onto the end of it, but with a slice you can
	slice = append(slice, 23)           // appends a slice

	// map
	// kind of like a dictionary - formula: variable = make(map[KeyValue]ValueType)
	sides := make(map[string]int)

	sides["triangle"] = 3
	sides["square"] = 4
	sides["megagon"] = 1000000 // a megagon is a shape with 1,000,000 sides lol

	// you can also make a map and define its elements 
	// in curly braces like this:

    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }


	fmt.Println(floats)			 // output: map[first:35.98 second:26.99]
	fmt.Println(sides)           // output: map[megagon:1000000 square:4 triangle:3]
	fmt.Println(sides["square"]) // get value of particular key

	// delete(sides, "triangle") // deletes element from map

	// for loop

	for i := 0; i < 10; i++ {
		sum += 1
	}

	/* for loop and while loop

	for loop

	for i:= 0; i > 5; i++{
		fmt.Println("I is now:", i)
	}

	while loop:

	for i < 5{
		fmt.Println("I is now:", i)
		i++
	}

	both output:
	1
	2
	3
	4
	5


	! look at this cool thing haha

	arr := []string{"a", "b", "c"}

	for index, value := range arr { // i think range gets index of element idk lol
		fmt.Println("index:" index, "value", value)

		output:
		index: 0 value: a
		index: 1 value: b
		index: 2 value: c
	}

	*/

	// cases:

	var day = 4  // what case will activate, e.g this has a value of 4 so only case(4) can activate. the value can be changed of course
	switch day { // in this example there is no case(4)
	case (1):
		fmt.Print("Saturday")
	case (2):
		fmt.Print("Sunday")
	default:
		fmt.Print("Weekday\n\n")
	}

	// see memory address of var
	fmt.Println(num1)  // 22
	fmt.Println(&num1) // 0x14000110018

	// pointers:
	// holds memory address of another var

	// nvm i dont know enough yet so...
	// here's how to find type (not value) of variable using printf 💀
	fmt.Printf("\nThe variable num1 is an %T\n", num1) // with %T you can output the type of a variable

	fmt.Println(cars, cars2, apples)
	fmt.Println(x)
	fmt.Println("Hello")
	fmt.Println(calc(num1, num2)) // 22, 5 would also work as parameter
	fmt.Println(sum)              // 45
	// fmt.Println(arr2)
}

// creating an array with the := operator only works while in a function
// array3 := [...]int{1, 2, 46}  <-- will not work
/*
func main() {
	array3 := [...]int{1, 2, 46}  <-- will work
}
*/var array3 = [...]int{1, 2, 46} // whereas with the var keyword and without the infer operator it is possible to make arrays outside of functions

// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

/*		EXERCISES:

! 21/11/23

1. Print i as long as i is less than 6:

func main() {
   for i:=0; i < 6; i++ {  // for [is is 0]; [while i smaller than 6]; [i + 1]
   	fmt.Println(i)
  }
}


2. Use a for loop to print "Yes" 5 times:

func main() {
   for i:=0; i < 5; i++{
    fmt.Println("Yes") // "Yes Yes Yes Yes Yes"
  }
}


3. Stop if the for loops is 5:

func main() {
  for i:=0; i < 10; i++ {
    if i==5 {
	 break
    }
  }
}

4. Create a function named myFunction and call it inside main():

func myFunction(){
	fmt.Println("My function has just been called!")
}

func main(){
	myFunction()
}

4.5. Call it twice:

func myFunction(){
	fmt.Println("My function has just been called!")
}

func main(){
	myFunction()
	myFunction()
}

5. Create a function and return the addition of 5 + x:

func myFunction(x int) int {
   return 5 + x
}

func main() {
  fmt.Println(myFunction(3))
}

6. Create a function that greets you with your name:

func familyName(name string) string{
	fmt.Println("Hello,", name, "nice to meet you!")
}

func main(){
	familyName("Daniel")
	familyName("Jean")
	familyName("Milo")
}
*/
