package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":6378")

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		conn.Write(([]byte("+OK\r\n")))
	}

}
