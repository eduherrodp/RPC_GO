package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	A, B float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Print("Select an operation (+, -, *, /): ")
		scanner.Scan()
		operation := scanner.Text()

		if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
			fmt.Println("Invalid operation")
			continue
		}

		fmt.Print("Enter the first number: ")
		scanner.Scan()
		// This will be work with negative numbers

		a, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Error parsing the first number", err)
			continue
		}

		fmt.Print("Enter the second number: ")
		scanner.Scan()
		b, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Error parsing the second number", err)
			continue
		}

		var result float64
		var args = Args{a, b}

		switch operation {
		case "+":
			err = client.Call("Calculator.Add", args, &result)
		case "-":
			err = client.Call("Calculator.Subtract", args, &result)
		case "*":
			err = client.Call("Calculator.Multiply", args, &result)
		case "/":
			err = client.Call("Calculator.Divide", args, &result)
		default:
			fmt.Println("Invalid operation")
			continue
		}
		if err != nil {
			fmt.Println("Error performing operation:", err)
			continue
		}
		fmt.Printf("Result: %.2f\n\n", result)
	}
}
