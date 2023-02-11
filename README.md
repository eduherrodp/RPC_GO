# RPC_GO
A simple calculator using Go's `net/rpc` package to perform operations over a network connection.

## Running the project
1. Clone the repository `git clone https://github.com/eduherrodp/RPC_GO.git` and `cd` into the directory.
2. In two different terminals `cd` into `client` and `server` , run `go run main.go` and `go run main.go` respectively.
3. Execute the server first, then the client.
4. The client will prompt for an operation and two numbers `(+, -, *, /)`, input your desired calculation and press enter.
5. The server will perform the operation and return the result to the client.
6. The client will print the result and prompt for a new operation.

## Notes
The server has to be running before the client can connect to it. The client will try to connect to the server every 5 seconds until it is able to connect.

## Author

[José Eduardo Hernández Rodríguez](https://github.com/eduherrodp).
