package main

import (
	"io"
	"log"
	"net"
	"os"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	logFatal(err)

	_, err = io.Copy(os.Stdout, conn)
	logFatal(err)
}
