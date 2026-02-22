package main

import "fmt"

func ds_main(){
	arrays()
	maps()
}

func maps(){
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":2, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Adam"] // second value is boolen whih is returned true if key actually exists else it returns false and a default value is returned 
	fmt.Println(age, ok)
	delete(myMap2,"Adam")

	for name, val:=range myMap2{
		fmt.Println(name, val)
	}
}

func arrays() {
	// var intArr [4]int32 =[4]int32{1,2,3}
	intArr := [...]int32{1,2,3} // still size 3
	fmt.Println(intArr[2])
	var intSlice []int32 = []int32{1,2,3}
	var intSlice2 []int32 = []int32{4,5,6}
	
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))
	
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))

	intSlice = append(intSlice, intSlice2...) // Appending multiple together
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 4, 9) 
	fmt.Println(intSlice3)

	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:])
	// println(&intArr)
	// fmt.Println(&intArr[0]) // Address 
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])
}