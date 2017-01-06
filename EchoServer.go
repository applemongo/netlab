package netlab

import (
	"net"
	"log"
	"fmt"
)

var serverPort = 50002
var serverBindIp = "0.0.0.0"
var laddr string = fmt.Sprint("%s:%s", serverBindIp, serverPort)

func Start(){
	ln, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalf("%s", err)
		return
	}

	fmt.Println(laddr)
	fmt.Println(ln)

	// Accept incoming socket connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("%s", err)
		}
        go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn){
    defer conn.Close()

    fmt.Println("remote address: %s", conn.RemoteAddr())
    //bufLen := 1024
    buf := make([]byte, 0)
    //buf := []byte{}

    for {
        dataBytes := make([]byte, 64)
        n, err := conn.Read(dataBytes)

        // We have an error, close socket and return
        if err != nil {
            log.Printf("fail to read data bytes: %v", err)
            break
        }
        buf = append(buf, dataBytes[:n]...)
        len := len(buf)

        log.Printf("buf length: %s", len)

        n, err = conn.Write(buf)

        // We have an error, close socket and return
        if err != nil {
            log.Printf("fail to write data bytes: %v", err)
            break
        }
        buf = buf[n:]
        len = len(buf)
    }


}