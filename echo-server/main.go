package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err != nil && err != io.EOF {
			log.Println("Unexpected error: ", err)
			break
		}

		if err == io.EOF && size == 0 {
			log.Println("Client disconnected: ", err)
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		log.Println("Writting data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to wrtie data: ", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection: ", err)
		}
		go echo(conn)
	}
}
