// functions and modules

// function declaration, parameters , return

package main

import "fmt"

// func add(a ...int)int{
// 	total :=0
// 	for _, v:= range a{
// 		total = total + v
// 	} 
// 	return total
// }

func main(){
	// result := add(5,6,343,53)
	// fmt.Println(result)

	//anonymous function

// 	result := func(a,b int)int{
// 		return a+b
// 	}(4,5)
// 	fmt.Println(result)
// 

// result := func(a,b int) int{
// 	return a+b
// }

// fmt.Println(result(4,8))

fmt.Println("Line 1")

defer fmt.Println("line 2")
defer fmt.Println("line 3")

fmt.Println("line 4")



}