package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		for {
			var data []byte
			conn.Read(data)

			var buf bytes.Buffer
			buf.Write(data)
			myData := buf.String()
			if myData != "" {
				fmt.Println(buf.String())
			}

			resp := "+PONG\r\n"
			conn.Write([]byte(resp))
		}

	}
}
