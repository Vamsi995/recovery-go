package main

import "fmt"

// import (
// 	"fmt"
// )

// import "errors"

func main() {
	// fmt.Println("Hey everyone")

	// var test string = "hello"
	// fmt.Println(test)

	// const answer int = 5
	// fmt.Println(answer)

	// var quotient, remainder, err = intDivision(4, 0)
	// // if err != nil {
	// // 	fmt.Printf("%s", err.Error())
	// // } else if remainder == 0 {
	// // 	fmt.Printf("The result of the integer divison is %v", quotient)
	// // } else {
	// // 	fmt.Printf("Quotient: %v, Remainder %v", quotient, remainder)
	// // }
	// switch {
	// 	case err != nil:
	// 		fmt.Printf("%s", err.Error())
	// 	case remainder == 0:
	// 		fmt.Printf("The result of the integer divison is %v", quotient)
	// 	default:
	// 		fmt.Printf("Quotient: %v, Remainder %v", quotient, remainder)

	// }

	// Arrays
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [...]int32{1, 2, 3}
	// intArr[0] = 1
	// intArr[1] = 2
	// intArr[2] = 3
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// fmt.Println(intArr[1:])

	// Slices
	// var intSlice []int32 = []int32{1, 2, 3}
	// fmt.Print((intSlice))
	// fmt.Println(len(intSlice), cap(intSlice))

	// intSlice = append(intSlice, 6)
	// fmt.Print(intSlice)

	// fmt.Println(len(intSlice), cap(intSlice))

	// var intSlice1 []int32 = make([]int32, 3, 8)
	// fmt.Println(intSlice1)

	// Maps

	// var hashmap map[string]int32 = make(map[string]int32)
	// fmt.Println(hashmap)

	// var hashmap_init = map[string]int32{"Vamsi": 24, "Pathma": 23}
	// fmt.Println(hashmap_init["test"])

	// var age, ok = hashmap["test"]
	// fmt.Printf("Age: %v, Exists: %v", age, ok)
	// fmt.Println()

	// delete(hashmap_init, "Vamsi")
	// delete(hashmap_init, "Pathma")

	// fmt.Println(hashmap_init)

	// Loops

	// for name, age := range hashmap_init {
	// 	fmt.Printf("Name: %v, Age: %v \n", name, age)
	// }

	// var intSlice []int32 = []int32{1, 2, 3}
	// for i, value := range intSlice {
	// 	fmt.Printf("Index: %v, Value: %v \n", i, value)
	// }

	// Infinite loops

	// var i int = 0

	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Print(i)
	// 	i += 1
	// }

	// for i:=0;i<10;i++ {
	// 	fmt.Println(i)
	// }

	// Strings
	// Strings in Go are byte arrays -> list of uint8
	// var mystring1 = "😭resume"
	// fmt.Printf("%T", mystring1[0])

	// // Converting to runes -> unicode point numbers that represent the character -> list of int32
	// var mystring = []rune("😭resume")
	// fmt.Printf("%T", mystring)
	// fmt.Println(len(mystring))
	// // var strBuilder strings.Builder
	// for i, v := range mystring {
	// 	fmt.Printf("%d, %c \n", i, v)
	// }

	// Structs and Interfaces

	// var myEngine gasEngine = gasEngine{225, 12}
	// fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Println("testing this %s")

}

// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("cannot divide by zero")
// 		return 0, 0, err
// 	}
// 	var answer int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return answer, remainder, err

// }

// type gasEngine struct{
// 	mpg uint8
// 	gallons uint8
// }
