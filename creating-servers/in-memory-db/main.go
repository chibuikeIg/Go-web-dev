package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	//instructions
	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n"+
		"USE:\n"+
		"SET key value \n"+
		"GET key \n"+
		"DELETE key \n\n"+
		"EXAMPLE:\n"+
		"SET fav chocolate\n"+
		"GET fav \n\n\n")

	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) < 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
				continue
			}

			k := fs[1]
			v := fs[2]

			if len(fs) > 3 {

				for i := 0; i < len(fs); i++ {

					if i != 0 && i != 1 && i != 2 {
						v += " " + fs[i]
					}

				}

			}

			data[k] = v

		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMANDS")
		}
	}

}
