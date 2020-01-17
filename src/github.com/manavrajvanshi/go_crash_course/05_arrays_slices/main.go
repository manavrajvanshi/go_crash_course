package main

import "fmt"


func main(){

	var arr1 [3]int
	arr1[0] = 213123
	arr1[1] = 42323423
	arr1[2] = 5234

	arr2 := [5]int{1,2,3,5}
	fmt.Println(arr1)
	fmt.Println(arr2)


	slice1 := []string{"ab", "cd", "ef","gh"}

	fmt.Println(slice1)

}
