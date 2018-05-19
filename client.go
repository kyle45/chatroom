package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var exit = make(chan bool)

func main() {
	log.Println("Before net.Dial")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	log.Println("After net.Dial")
	CheckError(err)
	defer conn.Close()
	log.Println("Before SendMsg")
	go SendMsg(conn)
	log.Println("After SendMsg")

	log.Println("Before RecvMsg")
	go RecvMsg(conn)
	log.Println("After RecvMsg")

	<-exit

}
func RecvMsg(conn net.Conn) {
	log.Println("Enter RecvMsg")
	var buf = make([]byte, 1024)
	for {
		log.Println("Before conn.Read")
		NumOfByte, err := conn.Read(buf)
		log.Println("After conn.Read")
		CheckError(err)
		fmt.Printf("Get Msg %s\n", string(buf[0:NumOfByte]))
	}

}
func SendMsg(conn net.Conn) {
	log.Println("Enter SendMsg")
	for {
		var input string
		log.Println("Before Input Msg")
		//fmt.Scanf("%s", &input)
		inputReader := bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		fmt.Println(input)
		log.Println("After Input Msg")

		if strings.ToUpper(input) == "EXIT" {
			log.Println("Input EXIT")
			exit <- true
			break
		}
		log.Println("Before conn.Write")
		conn.Write([]byte(input))
		log.Println("After conn.Write")

	}

}
func CheckError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
