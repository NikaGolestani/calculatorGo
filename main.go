package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Get value 1
	fmt.Println("Enter value 1:")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Nooooooooo You have to enter a number ")
	}

	//Get value 2
	fmt.Println("Enter value 2:")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Nooooooooo You have to enter a number ")
	}
	i := 0

	for i != 1 {
		fmt.Println("What should I do:(+ - *  /)")
		input3, _ := reader.ReadString('\n')
		equation1 := strings.TrimSpace(input3)
		//Get the sum
		switch equation1 {
		case "+":
			result := math.Round((float1+float2)*100) / 100
			fmt.Println(result)
			i = 1
		case "-":
			result := math.Round((float1-float2)*100) / 100
			fmt.Println(result)
			i = 1
		case "*":
			result := math.Round((float1*float2)*100) / 100
			fmt.Println(result)
			i = 1
		case "/":
			result := math.Round((float1/float2)*100) / 100
			fmt.Println(result)
			i = 1
		default:
			fmt.Println("Enter a valid operation!!!")
		}

	}
}
