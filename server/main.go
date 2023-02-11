package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Calculator is the structure that implements the service
type Calculator float64

type Args struct {
	A, B float64
}

// Add method
func (c *Calculator) Add(args *Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

// Subtract method
func (c *Calculator) Subtract(args *Args, reply *float64) error {
	// Thiw will be work with negative numbers
	if args.A < 0 && args.B < 0 {
		args.A, args.B = -args.A, -args.B
		*reply = -(args.A - args.B)
		return nil
	}
	if args.A < 0 {
		args.A, args.B = -args.A, -args.B
		*reply = -(args.A + args.B)
		return nil
	}
	if args.B < 0 {
		*reply = args.A + -args.B
		return nil
	}
	*reply = args.A - args.B
	return nil
}

// Multiply method
func (c *Calculator) Multiply(args *Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

// Divide method
func (c *Calculator) Divide(args *Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	err := rpc.Register(calculator)
	if err != nil {
		return
	}
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
