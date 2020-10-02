package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := c.SetDeadline(time.Now().Add(4 * time.Second))
	if err != nil {
		log.Println("connection timedout")
	}
	scanner := bufio.NewScanner(c) //c implements the Read()

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		fmt.Fprintf(c, "eco %s\n", text)
	}
	defer c.Close()
	fmt.Println("dot dot dot")
}
