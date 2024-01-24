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
		size, err := conn.Read(b)
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected error")
			break
		}

		log.Printf("Received %d bytes: %s", size, string(b))

		if _, err := conn.Write(b[:size]); err != nil {
			log.Fatal("Error during wright")
		}
	}
}

func main() {
	//Привязываемся к TCP порту, запускается на всех интерфейсах
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal("Unable to bind to port")
	}

	log.Println("Listening on port 0.0.0.0:20080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Println("Error during receiving conn")
		}
		go echo(conn)
	}
}
