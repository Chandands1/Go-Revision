
// Variables, Types & Core Expressions

package main

import "fmt"

const Pi = 3.14
const(
	A = iota
	B
	C 
	D
)

func main(){

	// //var declarations
	// var integer int
	// var word string
	// var flag bool
	// var number int = 45

	// fmt.Println(integer, string(word), flag, number)

	// //Short := declaration

	// value := 54
	// name := "Rahul"

	// fmt.Println(value, name)

	// // Basic types: bool, int/uint families, float64, complex, string

	// var bol bool = true
	// var number1 int = -1234
	// var number2 uint = 1234
	// var address string= "India"
	// var amount float64 = 64.34
	// var com complex64 = 43+9i

	// fmt.Println(bol, number1, number2, address, amount, com)

	// // byte vs rune

	// var by byte = 'A'
	// var ru rune = 'A'

	// fmt.Println(by, ru)

	// fmt.Println(Pi)

	// fmt.Println(A, B,C,D)


	//Type inference & explicit conversions

	// h := 5 // type inference
	// var k int = 5 //explicit declaration
	// var b float64 = float64(k) // explicit type conversion
	// fmt.Println(h,k)
	// fmt.Printf("%T", b)

	// =====Operations==========

	// Arthmatic Operations

	/*
	  + Addition
	  - Subtraction
	  * Multiplication
	  / Division
	  % Modulus
	*/

	// a := 10
	// b := 4

	// addition := a +b
	// substraction := a -b
	// multiplication := a*b
	// division := a/b
    // modulus := a%b

	// fmt.Println(addition, substraction, multiplication, division, modulus)


	// // comparison operators
	// /*
	//    it will return the bool value true or flase
	//    <
	//    <=
	//    >
	//    >=
	//    ==
	//    !=
	// */
    //  fmt.Println(a<b)
	//   fmt.Println(a<=b)
	//    fmt.Println(a>b)
	//     fmt.Println(a>=b)
	// 	 fmt.Println(a==b)
	// 	  fmt.Println(a!=b)

	// if a <= 10 && b <= 10{
	// 	fmt.Println("Numbers are less than or equal to 10", a, b)
	// }
	// if a <10 || b <10{
	// 	fmt.Println("Any number is less than 10")
	// }

	//Scope rules, shadowing, blank identifier _


	simpleCalculator()


}

func simpleCalculator() {
	var choose int

	fmt.Println("Choose the operation\n1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n5.Modulus")
	fmt.Scanln(&choose)

	result := menu(choose)
	fmt.Println("Result:", result)
}

func menu(choose int) int {
	var number1, number2 int

	fmt.Println("Enter number1:")
	fmt.Scanln(&number1)

	fmt.Println("Enter number2:")
	fmt.Scanln(&number2)

	switch choose {
	case 1:
		return number1 + number2
	case 2:
		return number1 - number2
	case 3:
		return number1 * number2
	case 4:
		if number2 == 0 {
			fmt.Println("Cannot divide by zero")
			return 0
		}
		return number1 / number2
	case 5:
		return number1 % number2
	default:
		fmt.Println("Invalid choice")
		return 0
	}
}