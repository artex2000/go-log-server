package main

import (
	"github.com/Microsoft/go-winio"
	"log"
	"net"
	"os"
	"io"
)

func main() {
	listener, err := winio.ListenPipe("\\\\.\\pipe\\console_log", nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer listener.Close()

	log.Println("[server]: Session started *************")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[server]: Error accepting connection")
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil && err != io.EOF {
			log.Println("[server]: Error reading connection")
			return
		}
		_, err = os.Stdout.Write(buf[0:n])
		if err != nil {
			log.Println("[server]: Error writing to Stdout")
			return
		}
	}
}
