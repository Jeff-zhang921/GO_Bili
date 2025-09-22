package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//after connect need close
	defer conn.Close()

	for {
		//put the data to slice
		buf := make([]byte, 1024)
		//read the from connection
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[0:n]))
	}
}
func main() {
	fmt.Println("server start..")
	//use listen function
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("server listenerror:", err)
		return
	}
	//cycle
	for {
		conn, err2 := listen.Accept()

		if err2 != nil {
			fmt.Println("wait for customer fail:", err2)
			return

		} else {
			fmt.Println("customer connect success, conn:", conn, "getting customer info:", conn.RemoteAddr().String())
		}
		go process(conn)
	}

}
