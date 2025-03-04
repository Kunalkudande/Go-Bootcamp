package main

import (
	"fmt"
)

func main(){

//1) variable and datatypes
	// Explicit declaration
	var ticket int = 50 
	var name string

	// Type inference (shortcut)
    age := 30  // int type

	// Constants
    const pi = 3.14

	//Basic Types: int, float64, bool, string, array, slice, map, struct
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)
	fmt.Printf("ticket type is %T \n", ticket)    //%T=type, %s=string, %d=value, %q=for quote like "Kunal", %p=pointer address 
	fmt.Println("Your age is: ", age)

//2) If Else = after IF complete then with same line else should start
	num := 12

	if num < 0 {
		fmt.Println("Negative number")
	} else if num > 18 {
		fmt.Println("your are younger")
	} else{
		fmt.Println("below 18")
	}

//3) For Loop = Go has only for (no while):
	// Classic for loop
	for i := 0; i < 3; i++ {
		fmt.Print(i)  // 0, 1, 2
		fmt.Print(" ")
	}
	fmt.Println()
	// While-like loop
	count := 0
	for count < 2 {
		fmt.Println(count)  // 0, 1
		count++
	}

	// Range loop (slice)
    nums := []int{10, 20, 30}
    for idx, val := range nums {
        fmt.Printf("Range Loop - Index: %d, Value: %d\n", idx, val)
    }

//4) switch case 
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Workday")
	case "Tuesday":
		fmt.Println("Taco Tuesday!")  // Output
	default:
		fmt.Println("Other day")
	}

	//or
	number := 2
	switch number {
	case 1:
		fmt.Println("Hi I am 1")
	case 2:
		fmt.Println("Hi I am 2")
	default:
		fmt.Println("Enter vaild number")	
	}
//5) functions
	// Basic function
	func add(a int, b int) int {
		return a + b
	}

	// Multiple return values
	func swap(x, y string) (string, string) {
		return y, x
	}

	sum := add(3, 5)  // 8
	a, b := swap("hello", "world")
	fmt.Println(sum, a, b)  // 8 world hello

//6) Arrays, Slices, Maps, struct
	//Array (fixed size)
	var numbers [3]int = [3]int{1, 2, 3}   //or numbers := [3]int{1, 2, 3}
	// Access elements
	fmt.Println("First Element:", numbers[0])
	fmt.Println("Last Element:", numbers[len(numbers)-1])

	//Slice (dynamic array)
	fruits := []string{"Apple", "Banana"}
	fruits = append(fruits, "Cherry")  // Add item
	fmt.Println(fruits[1:3])  // Slice from index 1 to 2: [Banana Cherry]

	//Map (key-value pairs)
	user := map[string]string{
		"name": "John",
		"age":  "30",
	}
	fmt.Println(user["name"])  // John

	//struct 
	type Person struct { // Struct
        FirstName string
        LastName  string
        Age       int
    }
    var person = Person{"Alice", "Smith", 30}

//7) Pointer
	x := 10
	ptr := &x  // Get address of x
	fmt.Println(*ptr)  // 10 (value at address)

	*ptr = 20  // Change x via pointer
	fmt.Println(x)  // 20
}