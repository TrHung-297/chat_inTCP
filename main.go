package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	go s.run()

	listener, err := net.Listen("tcp",":8888")
	if err!= nil {
		log.Fatalf("%v",err.Error())
	}
	defer listener.Close()

	log.Printf("start server on:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}