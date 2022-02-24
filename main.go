package main

import (
	"github.com/HasinduLanka/console"
)

func main() {

	// console.WriteToFile("out.txt")
	// console.ScanFromFile("in.txt")

	console.Print("Hashcode 2022")

	console.Print("Enter your name: ")
	s := console.ReadLine()
	console.Print("Hello, " + s)

	console.Print("\nEnter an array of strings seperated by spaces: ")
	A := console.ReadArrayClean(" ")
	console.Print("This is the array you entered: ")
	console.Log(A)

	console.Print("\nEnter an array of strings seperated by spaces to make a Hashset : ")
	H := console.ReadHashset(" ")
	console.Print("This is the array you entered as a Hashset : ")
	console.Log(H)

	console.Print("\nEnter key value pairs seperated by spaces : (Example - key1 value1 key2 value2 ) ")
	P := console.ReadPairs(" ", 0)
	console.Print("This is the Hashmap you entered : ")
	console.Log(P)
}
