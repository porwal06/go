package practice

import "fmt"

//  1. Create a new array (!) that contains three hobbies you have
//     Output (print) that array in the command line.
//  2. Also output more data about that array:
//     - The first element (standalone)
//     - The second and third element combined as a new list
//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.

// See example for map, slice, array with `make()` function & custom datatype.

type Product struct {
	title        string
	id           int
	price        float64
	product_list []string
}

// Create custom type
type productCustomMapType map[string]float64

func (s productCustomMapType) printMap() {
	fmt.Println(s)
}

func PracticeList() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	arrayHobby := [3]string{"Coding", "Gaming", "Reading"}

	fmt.Println("Task1 - Total list of hobbies:", arrayHobby, "\n")

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Printf("Task2 - First Hobby: %s Second & third Hobby: %s\n\n", arrayHobby[0], arrayHobby[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	sliceHobby := arrayHobby[0:2] // First way
	sliceHobby2 := arrayHobby[:2] // Second way

	fmt.Printf("Task3 - SlicedHobby1: %v, SliceHobby2: %v\n\n", sliceHobby, sliceHobby2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	resliceHobby := sliceHobby[1:3]
	//sliceHobby = append(sliceHobby, "Sleeping")

	fmt.Println("Task4 - Contain 2nd & last item of original array", resliceHobby, "\n")

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	arrGoals := []string{"Learn Go", "Learn Symfony"}
	fmt.Printf("Task5 - Goal array: %v\n\n", arrGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	arrGoals[1] = "Learn React"
	arrGoals = append(arrGoals, "Learn Angular")
	fmt.Println("Task6 - Updated Goal array: ", arrGoals, "\n")

	//  7. Bonus: Create a "Product" struct with title, id, price and create a
	//     dynamic list of products (at least 2 products).
	//     Then add a third product to the existing list of products.

	product_struct := []Product{
		{
			title:        "Product 1",
			id:           1,
			price:        9.99,
			product_list: []string{"Product 11", "Product 12"},
		},
		{
			title:        "Product 2",
			id:           2,
			price:        19.99,
			product_list: []string{"Product 21", "Product 22"},
		},
	}

	fmt.Println("Task7\n - Product struct:", product_struct)

	product_struct = append(product_struct, Product{
		title:        "Product 3",
		id:           3,
		price:        39.99,
		product_list: []string{"Product 31", "Product 32"},
	},
	)
	fmt.Println(" - Product struct with 3 product list:", product_struct, "\n")

	MapExample()

	SliceExample()

}

// Create dynamic array for Product struct

// array example
func ListArray() {
	var arr [4]int = [4]int{6, 7, 9, 10}

	arr2 := [3]int{6, 7, 9}

	fmt.Println(arr)         //[6 7 9 10]
	fmt.Printf("%v\n", arr2) //[6 7 9]

	//Slice
	highPrice := arr[1:3]
	fmt.Println(highPrice) //[7 9]
}

// Map example
// It is similar to array/slice but inplace of index we can define any key
func MapExample() {

	websites := map[string]string{
		"Google":   "google.com",
		"Facebook": "facebook.com",
		"Twitter":  "x.com",
	}
	websites["Instagram"] = "instagram.com"
	fmt.Println("MapExample:", websites)
	products := productCustomMapType{
		"Keyboard": 99.99,
		"Mouse":    59.99,
	}
	products.printMap()

}
func SliceExample() {
	//Slice
	slice1 := make([]int, 3)
	slice1[0] = 45

	fmt.Println("Slice Example:\n", slice1)
	slice2 := []int{4, 5, 6, 7, 8}
	fmt.Println(slice2)

	//appending 2 slice
	mergeSlice := append(slice1, slice2...)
	fmt.Println("Slice exaple using make function: ", mergeSlice, "\n")
}
