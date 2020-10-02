package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to start server", err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	scanner := bufio.NewScanner(c) //c implements the Read()

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	defer c.Close()
	fmt.Println("dot dot dot")
}
