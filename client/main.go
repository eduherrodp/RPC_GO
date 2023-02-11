package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type Argumentos struct {
	A, B int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Print("Seleccione una operación (+, -, *, /): ")
		scanner.Scan()
		operacion := scanner.Text()

		fmt.Print("Ingrese el primer número: ")
		scanner.Scan()
		a, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		fmt.Print("Ingrese el segundo número: ")
		scanner.Scan()
		b, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		var result int
		var args = Argumentos{a, b}

		switch operacion {
		case "+":
			err = client.Call("Calculadora.Sumar", args, &result)
		case "-":
			err = client.Call("Calculadora.Restar", args, &result)
		case "*":
			err = client.Call("Calculadora.Multiplicar", args, &result)
		case "/":
			err = client.Call("Calculadora.Dividir", args, &result)
		default:
			fmt.Println("Operación inválida")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Resultado: %d\n\n", result)
	}
}
