package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			ln := scanner.Text()

			fmt.Println(ln)

			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP HEADER REQUEST")
				break
			}
		}

		fmt.Println("Code got here")

		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
