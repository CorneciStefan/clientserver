package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("Enter text:")
		text, err2 := stdin.ReadString('\n')
		if err2 != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(conn, text)
		read(&conn)
	}
}
