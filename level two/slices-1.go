package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{
		1, 2, 3,
	}

	fmt.Println(arr1)

	arr1 = append(arr1, 4)

	fmt.Println(arr1)

	arr1 = append(arr1, 1, 2, 3, 4)

	fmt.Println(arr1)

	arr2 := []int{}

	fmt.Println(arr2)

	fmt.Println(len(arr2))

	arr2 = append(arr2, arr1...)
	
	arr2 = append(arr2,arr1...);
	
	fmt.Println(arr2)
	
	sort.Ints(arr2)

	fmt.Println(arr2);
	
	letter := []byte{'l','o','v','u'}
	
	fmt.Println(letter[0:][0]=='l');
	
}
