package main

import (
	"container/list"
	"fmt"
	"net"
)

var conn_list list.List

func main() {
	// todo create tcp server
	tcpServer, _ := net.ResolveTCPAddr("tcp4", ":8080")
	listener, _ := net.ListenTCP("tcp", tcpServer)

	// accpet scoket loop
	for {
		conn, err := listener.Accept()
		if isError(err) {
			continue
		}
		conn_list.PushBack(conn)
		a := conn_list.Len()
		fmt.Println("len is ", a)
		go handle(conn)
	}
}

// todo every socket recv the message and send all client in list
func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("连接加一，客户端地址是:", conn.RemoteAddr())
	a := "i am go mess"
	conn.Write([]byte(a))

	buf := make([]byte, 1024)
	for {
		// recv the message
		len, err := conn.Read(buf)
		// if it have error in programming , break and return !!!
		if isError(err) {
			break
		}

		// todo sned message to all client socket
		send_all(string(buf))
		fmt.Println("clent mess is ", string(buf), "------", "length is ", len)

	}
}

func send_all(mess string) {
	for i := conn_list.Front(); i != nil; i = i.Next() {
		// element trans to net.conn
		conn, ok := (i.Value).(net.Conn)
		if !ok {
			continue
		}
		conn.Write([]byte(mess))

	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
