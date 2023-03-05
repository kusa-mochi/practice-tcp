package main

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/eiannone/keyboard"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()

	for {
		// キーボード入力待機
		_, _, err := keyboard.GetSingleKey()
		logFatal(err)
		_, err = io.WriteString(conn, "Socket Connection!!\n")
		logFatal(err)
		// time.Sleep(time.Second)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	logFatal(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	logFatal(err)

	err = keyboard.Open()
	logFatal(err)
	defer keyboard.Close()

	for {
		// err := listener.SetDeadline(time.Now().Add(time.Second * 10))
		err := listener.SetDeadline(time.Time{})
		logFatal(err)

		conn, err := listener.AcceptTCP()
		logFatal(err)

		go echoHandler(conn)
	}
}
