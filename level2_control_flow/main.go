//Control Flow

package main

import "fmt"

func main(){

	// if, else if, if else

	// age := 19
	// if age < 18{
	// 	fmt.Println("Not adult")
	// }else{
	// 	fmt.Println("Adult")
	// }

	// day := 3
	// if day == 1{
	// 	fmt.Println("Monday")
	// }
	// if day == 2{
	// 	fmt.Println("Tuesday")
	// }
	// if day ==3{
	// 	fmt.Println("Wednesday")
	// }

	// marks := 88

	// if marks >=90{
	// 	fmt.Println("A")
	// }else if marks >= 80{
	// 	fmt.Println("B")
	// }

	//if with short initialization

	// if age:= 17; age < 18{
	// 	fmt.Println("Not adult")
	// }

	//switch (expression, type, fallthrough)

	// day := 5

	// switch day{
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4 , 5:
	// 	fmt.Println("Thursday")
	// default:
	// 	fmt.Println("Others day")
	// }

	// x := 6

	// switch {

	// case x < 10:
	// 	fmt.Println("number is less than 10")
	// 	fallthrough
	// case x >10:
	// 	fmt.Println("number is greater than 10")
	// case x ==10:
	// 	fmt.Println("number is equal to 10")
	// default:
	// 	fmt.Println("Noting matches")

	// }


	//for (classic 3-part, condition-only, infinite)

	// for i:=0; i <= 10;i++{
	// 	fmt.Println(i)
	// }
	// i:=4

	// for i <=10{
	// 	fmt.Println("Hello")
	// 	i++
	// }

	// for{
	// 	fmt.Println("Infinite")
	// }

	//range over slices/maps/strings


	// sli := []int{21,23,45,543,664}
	//  m := map[string]string{
	// 	"1":"A",
	// 	"2":"B",
	// 	"3":"C",}
	//  for k,v := range m{
	// 	fmt.Println(k,v)
	//  }

	
	//  for i,v := range sli{
	// 	fmt.Println(i,v)
	// }

	// s := "hello word"
	// for i, v := range s{
	// 	if string(v) == " "{
	// 		continue

	// 	}
	// 	fmt.Println(i,string(v))
	// }

	// break and continue

	// for i:=0 ; i <=10 ;i++{
	// 	// if i ==2{
	// 	// 	break
	// 	// }
	// 	if i == 2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }


}