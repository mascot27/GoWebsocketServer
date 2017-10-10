package main

import (
	"fmt"
)


func ParseMessage(message []byte)([]string){

	part1 := ""
	part2 := ""
	part3 := ""

	fmt.Println("---start---")
	for i:=0 ; i<len(message) ; i++ {
		fmt.Println(string(message[i]))
	}

	fmt.Println("---end---")

	myStringArray := []string{part1, part2, part3}
	return myStringArray

}
