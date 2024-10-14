package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		panic("dial err: " + err.Error())
	}
	defer conn.Close()

	for {
		data := make([]byte, 1024)
		_, err := conn.Read(data)
		if err != nil {
			if err != io.EOF {
				panic("read err: " + err.Error())
			}

			break
		}

		fmt.Printf(string(data))
	}

	fmt.Println()

}
