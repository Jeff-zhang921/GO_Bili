package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//different computer connecting, sending data
//IP+PORT
//Ip is an address, each computer has it own IP
//port is in computer, one computer has many port
//each app has a port

//TCP/IP protocal two machine, one is pack one is unpack the data
//3 time shake hand principle to connect
//4 time shake hand to disconnect

// create customer
// using Dial function
func main() {
	fmt.Println("customer start..")
	//para: the protocal and server IP+port number
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("customer connect fail:", err)
		return
	}
	fmt.Println("customer connect success, conn:", conn)
	//os.stdin is the standard input of the terminal
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	//str send to server
	n, _ := conn.Write([]byte(str))
	fmt.Println("terminal data send success:", n, "of byte sent")
}
