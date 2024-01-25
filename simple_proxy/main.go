package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "localhost:20080")
	if err != nil {
		log.Fatalln("Unable to connect to our host")
	}

	defer dst.Close()

	//Выполняем в горутине для предотвращения блокировки
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// прослушиваем локальный порт
	listener, err := net.Listen("tcp", ":80")

	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}
