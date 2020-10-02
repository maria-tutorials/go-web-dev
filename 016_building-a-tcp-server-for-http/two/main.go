package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var body string = `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
<strong>INDEX</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
</body></html>`

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	request(c)
}

func request(c net.Conn) {
	i := 0
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// parse/read request line
			m := strings.Fields(ln)[0] //method
			u := strings.Fields(ln)[1] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)

			mux(c, m, u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(c net.Conn, m string, u string) {
	if m == "GET" && u == "/" {
		index(c)
	} else if m == "GET" && u == "/about" {
		about(c)
	} else {
		notFound(c)
	}
}

func index(c net.Conn) {
	body := body
	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}

func about(c net.Conn) {
	body := body
	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}

func notFound(c net.Conn) {
	fmt.Fprint(c, "HTTP/1.1 404 NOT FOUND\r\n")
}
