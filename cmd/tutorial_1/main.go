package main

import (
	"errors"
	"fmt"
)

func main(){

	fmt.Println("Hello World!")
	// var intNum int 
	// fmt.Println(intNum)
	// myVar := "text"
	// fmt.Println(myVar)
	// var1, var2 := 1,2
	// fmt.Println(var1, var2)
	var inputParam string = "PrintString"
	printMe(inputParam)
	
	numerator, denominator := 10, 0
	value, value2, err := intDivision(numerator, denominator)
	// if(err==nil){
	// 	printInt(value, value2)
	// }else{ 
	// 	fmt.Printf(err.Error())
	// }
	switch(err){
	case nil:
		printInt(value, value2)
		// break this is implied in GO
	// case 1, 2:
	// 	do something 
	default:
		fmt.Printf(err.Error())

	}
}

func printInt(value int, value2 int){
	fmt.Printf("The division of the number is %v and remainder is %v ",value, value2)
}

func printMe(printName string){
	fmt.Println(printName)
}

func intDivision(num int, den int) (int, int, error){
	var err error
	if den==0{
		err = errors.New("Denominator cannot be 0")
		return 0,0,err
	}
	var res int = num/den
	var res2 int = num%den
	return res, res2, err
}