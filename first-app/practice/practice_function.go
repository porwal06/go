/**
* Function practice
*  Parameter function
*  Return function
*  Anonymous function
*  Closure
*  Recursion
*  Variadic function
 */

package practice

import (
	"fmt"
)

func PracticeFunction() {

	parameterFunction()
	returnFunction()
	anonymousFunction()
	closureFunction()
	recursionFunction()
	variadicFunction(2, 3, 4, 5, 6)
}

func parameterFunction() {
	numberSlice := []int{6, 7, 8, 9, 10}
	fmt.Println("Parameter function example", numberSlice)
	doubled := transformedNumber(&numberSlice, double)
	trippled := transformedNumber(&numberSlice, tripple)
	fmt.Println("Doubled number", doubled)
	fmt.Println("Trippled number", trippled)
}

func returnFunction() {
	numberSlice1 := []int{2, 7, 8, 9, 10}
	numberSlice2 := []int{3, 7, 8, 9, 10}
	doubled := transformedNumber(&numberSlice1, getTransformFunction(&numberSlice1))
	trippled := transformedNumber(&numberSlice2, getTransformFunction(&numberSlice2))
	fmt.Println("\nReturn function example", numberSlice1, numberSlice2)

	fmt.Println("Doubled number", doubled)
	fmt.Println("Trippled number", trippled)

}

func anonymousFunction() {
	numberSlice := []int{6, 7, 8, 9, 10}
	fmt.Println("\nAnonymous function example", numberSlice)
	doubled := transformedNumber(&numberSlice, func(val int) int {
		return val * 2
	})
	trippled := transformedNumber(&numberSlice, func(val int) int {
		return val * 3
	})
	fmt.Println("Doubled number", doubled)
	fmt.Println("Trippled number", trippled)
}

func closureFunction() {
	numberSlice := []int{6, 7, 8, 9, 10}
	fmt.Println("\nClosure function example", numberSlice)

	doubled := transformedNumber(&numberSlice, createFactor(2))
	trippled := transformedNumber(&numberSlice, createFactor(3))

	fmt.Println("Doubled number", doubled)
	fmt.Println("Trippled number", trippled)
}

// This is example for recursion
func recursionFunction() {
	squared := recursionFunctionHelper(5)
	fmt.Println("\nRecursion function example ie. squared of 5 - ", squared)
}

func variadicFunction(numberSlice ...int) {
	fmt.Println("\nVariadic function example", numberSlice)
	doubled := transformedNumber(&numberSlice, createFactor(2))
	trippled := transformedNumber(&numberSlice, createFactor(3))

	fmt.Println("Doubled number", doubled)
	fmt.Println("Trippled number", trippled)

}

// This is example for closure function
func createFactor(factor int) func(int) int {
	return func(val int) int {
		return val * factor
	}
}

// Parameter function "getTransformFunction" - Used in Return function
// Retrun function "double or tripple" - Used in Return function
func getTransformFunction(number *[]int) func(int) int {
	if (*number)[0] == 2 { // For getting value from pointer index number variable came within bracker else it give error.
		return double
	}
	return tripple
}

// Helper function Get function as parameter check transFn - Used in Parameter, retrun, anonymous, closure function
func transformedNumber(number *[]int, transFn func(int) int) []int {
	//transFn := double
	transformed := make([]int, len(*number))
	for index, val := range *number {
		transformed[index] = transFn(val)
	}
	return transformed
}

// Doubled number for passed number
func double(val int) int {
	return val * 2
}

// Tripple number for passed number
func tripple(val int) int {
	return val * 3
}

func recursionFunctionHelper(number int) (squared int) {
	if number == 0 {
		return 1
	}
	return number * recursionFunctionHelper(number-1)
}
