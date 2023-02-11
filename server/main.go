package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Calculadora es la estructura que implementa el servicio
type Calculadora int

// Argumentos es la estructura que contiene los argumentos de la operación
type Argumentos struct {
	A, B int
}

// Sumar
func (c *Calculadora) Sumar(args *Argumentos, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// Restar
func (c *Calculadora) Restar(args *Argumentos, reply *int) error {
	*reply = args.A - args.B
	return nil
}

// Multiplicar
func (c *Calculadora) Multiplicar(args *Argumentos, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Dividir
func (c *Calculadora) Dividir(args *Argumentos, reply *int) error {
	if args.B == 0 {
		return errors.New("División por cero")
	}
	*reply = args.A / args.B
	return nil
}
func main() {
	calculadora := new(Calculadora)
	rpc.Register(calculadora)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
