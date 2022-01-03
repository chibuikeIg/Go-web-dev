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
		log.Panic(err)
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)
	}

	defer conn.Close()

	fmt.Println("Code got here")

}
