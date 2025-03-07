package main

import (
	"fmt"
)

// Story 1: Given a list of integers, write a program to return only the even numbers from this list.
func evenNumber() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Print(numbers[i])
			fmt.Print(" ")
		} else {
			continue
		}
	}
}

// Story 2: Given a list of integers, write a program to return only the odd numbers from this list.
func oddNumbers(numbers [10]int) []int {
	var output []int
	for _, n := range numbers {
		if n%2 != 0 {
			output = append(output, n)
		}
	}
	return output
}

// checking and return bool
func isPrime(n int) bool {
	flag := true
	if n == 0 || n == 1 {
		flag = false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flag = false
			break
		}
	}
	return flag
}

// Story 3: Given a list of integers, write a program to return only the prime numbers from this list.
func primeNumber() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numbers); i++ {
		if isPrime(numbers[i]) {
			fmt.Print(numbers[i])
			fmt.Print(" ")
		}
	}
}

// Story 4: Given a list of integers, write a program to return only the odd prime numbers from this list.
func primeNumberAndOddNumber(number [10]int) []int {
	var output []int
	for _, n := range number {
		if isPrime(n) && n%2 != 0 {
			output = append(output, n)
		}
	}
	return output
}

// Story 5: Given a list of integers, write a program to return only the even and multiples of 5 from this list.
func evenAndMultipleOf5(num []int) []int {
	var output []int
	for _, n := range num {
		if n%5 == 0 && n%2 == 0 {
			output = append(output, n)
		}
	}
	return output
}

// Story 6: Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
func oddMultipleOf3GreaterThan10(num []int) []int {
	var output []int
	for _, n := range num {
		if n > 10 && n%3 == 0 && n%2 != 0 {
			output = append(output, n)
		}
	}
	return output
}

/*Story 7: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ALL the conditions.*/

func filter(numbers []int, conditions ...func(int) bool) []int {
	var output []int

	for _, n := range numbers {
		match := true
		for _, condition := range conditions {
			if !condition(n) {
				match = false
				break
			}
		}
		if match {
			output = append(output, n)
		}
	}
	return output
}

func isOdd(n int) bool {
	return n%2 != 0
}
func isGreaterThan5(n int) bool {
	return n > 5
}
func isMultipleOf3(n int) bool {
	return n%3 == 0
}
func isEven(n int) bool {
	return n%2 == 0
}
func isLessThan15(n int) bool {
	return n < 15
}

// Story 8: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ANY of the conditions.
func filterAny(numbers []int, conditions ...func(int) bool) []int {
	var output []int

	for _, n := range numbers {
		match := false
		for _, condition := range conditions {
			if condition(n) {
				match = true
			}
		}
		if match {
			output = append(output, n)
		}
	}
	return output
}

func isGreaterThan15(n int) bool {
	return n > 15
}
func isMultipleOf5(n int) bool {
	return n%5 == 0
}
func isLessThan6(n int) bool {
	return n < 6
}

func main() {
	//1)
	//evenNumber()
	////expected output = [2, 4, 6, 8, 10]

	//2)
	// numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	//expected output = [1, 3, 5, 7, 9]
	// var out = oddNumbers(numbers)
	// fmt.Print(out)

	//3)
	//primeNumber()
	//expected output = [2, 3, 5, 7]

	//4)
	// numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	// var out = primeNumberAndOddNumber(numbers)
	// fmt.Print(out)
	//expected output = [3, 5, 7]

	//5)
	// num := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// var output = evenAndMultipleOf5(num);
	// fmt.Print(output)
	//expected output = [10, 20]

	//6)
	// num := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// output := oddMultipleOf3GreaterThan10(num)
	// fmt.Print(output)
	//expected output = [15]

	//7)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//7.1
	//Conditions specified using a set of functions: odd, greater than 5, multiple of 3
	// result := filter(num, isOdd, isGreaterThan5, isMultipleOf3)
	// fmt.Print(result)
	//expected output: 9, 15

	//7.2
	//Conditions specified using a set of functions: even, less than 15, multiple of 3
	//result := filter(num, isEven, isLessThan15, isMultipleOf3)
	//fmt.Print(result)
	//expected output: 6, 12

	//8)
	//8.1
	//Conditions specified using a set of functions: prime, greater than 15, multiple of 5
	// result := filterAny(num, isPrime, isGreaterThan15, isMultipleOf5)
	// fmt.Print(result)
	//expected output: 2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20

	//8.2
	//Conditions specified using a set of functions: less than 6, multiple of 3
	result := filterAny(num, isLessThan6, isMultipleOf3)
	fmt.Print(result)
	//expected output: 1, 2, 3, 4, 5, 6, 9, 12, 15, 18
}
