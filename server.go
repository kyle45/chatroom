package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

var onlineConn = make(map[string]net.Conn)

func main() {
	log.Println("Before net.Listen")
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		CheckError(err)

	} else {
		log.Println("After net.Listen")
	}
	defer listen_socket.Close()
	fmt.Println("Server is waiting...")
	//go ConsumeMessage()
	for {
		conn, err := listen_socket.Accept()
		log.Printf("new user login %s", conn.RemoteAddr().String())
		CheckError(err)
		onlineConn[conn.RemoteAddr().String()] = conn
		go ProcessCli(conn)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
}

func SendMsg(addr string, msg string) {
	log.Println("Enter SendMsg")
	conn := onlineConn[addr]
	log.Println("Before conn.Write")
	conn.Write([]byte(msg))
	log.Println("After conn.Write")

}
func RecvMsg(conn net.Conn) string {
	log.Println("Enter RecvMsg")
	var buf = make([]byte, 1024)
	log.Println("Before conn.Read")
	NumOfByte, err := conn.Read(buf)
	if err != nil {
		return ""
	}
	log.Println("After conn.Read")
	CheckError(err)
	fmt.Printf("Get Msg %s\n", string(buf[0:NumOfByte]))
	return string(buf[0:NumOfByte])
}
func ProcessCli(conn net.Conn) {
	// 处理消息
	for {
		log.Println("Enter ProcessCli")
		content := RecvMsg(conn)
		if content == "" {
			return
		}
		//分割地址与msg
		//fmt.Println(NumOfByte)

		log.Println(content)
		strs := strings.Split(content, "#")
		var addr, msg string

		if len(strs) >= 2 {
			addr = strings.Split(content, "#")[0]
			msg = strings.Split(content, "#")[1]
			log.Println("Before SendMsg")
			SendMsg(addr, msg)
			log.Println("After SendMsg")
		} else {
			SendMsg(conn.RemoteAddr().String(), "Msg Error")
		}

		//向指定地址发功msg

	}

}
