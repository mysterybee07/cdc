package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert int to float64
	var intVal int = 42
	var floatVal float64 = float64(intVal)
	fmt.Printf("int to float64: %d to %f\n", intVal, floatVal)

	// Convert float64 to int
	var anotherFloatVal float64 = 42.57
	var anotherIntVal int = int(anotherFloatVal)
	fmt.Printf("float64 to int: %f to %d\n", anotherFloatVal, anotherIntVal)

	// Convert int to string
	var intToString int = 123
	var stringFromInt string = strconv.Itoa(intToString)
	fmt.Printf("int to string: %d to %s\n", intToString, stringFromInt)

	// Convert string to int
	var stringToInt string = "456"
	intFromString, err := strconv.Atoi(stringToInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("string to int: %s to %d\n", stringToInt, intFromString)
	}

	// Convert float64 to string
	var floatToString float64 = 123.456
	var stringFromFloat string = strconv.FormatFloat(floatToString, 'f', -1, 64)
	fmt.Printf("float64 to string: %f to %s\n", floatToString, stringFromFloat)

	// Convert string to float64
	var stringToFloat string = "789.012"
	floatFromString, err := strconv.ParseFloat(stringToFloat, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	} else {
		fmt.Printf("string to float64: %s to %f\n", stringToFloat, floatFromString)
	}
}
