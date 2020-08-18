package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func client() {
	server := "127.0.0.1:8080"
	addr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err)

	conn, err := net.DialTCP("tcp4", nil, addr)
	checkError(err)

	_, err = conn.Write([]byte("hello i am best"))
	checkError(err)

	response, _ := ioutil.ReadAll(conn)
	fmt.Println(string(response))
	os.Exit(0)
}
