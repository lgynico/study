package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"testing"
	"time"
	"zinx/znet"
)

func TestClient(t *testing.T) {

	log.Println("Client Test start...")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Println("connect server fail:", err)
		return
	}
	defer conn.Close()

	for {

		message := znet.NewMessage(0, []byte("ZINX ping client"))
		packet := &znet.Packet{}
		data, err := packet.Pack(message)
		if err != nil {
			log.Println("pack message fail:", err)
			return
		}

		if _, err = conn.Write(data); err != nil {
			log.Println("send message fail:", err)
			return
		}

		buffer := make([]byte, packet.HeaderLen())
		if _, err := io.ReadFull(conn, buffer); err != nil {
			log.Println("read message fail:", err)
			return
		}

		message, err = packet.Unpack(buffer)
		if err != nil {
			log.Println("unpack message fail:", err)
			return
		}

		buffer = make([]byte, message.DataLen())
		if _, err := io.ReadFull(conn, buffer); err != nil {
			log.Println("read message fail:", err)
			return
		}

		reader := bytes.NewReader(buffer)
		data = make([]byte, message.DataLen())
		if err := binary.Read(reader, binary.LittleEndian, &data); err != nil {
			log.Println("read message fail:", err)
			return
		}

		log.Printf("server call back: %s [%d]\n", data, message.DataLen())

		time.Sleep(time.Second)
	}
}

func TestClient1(t *testing.T) {

	log.Println("Client Test start...")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Println("connect server fail:", err)
		return
	}
	defer conn.Close()

	for {

		message := znet.NewMessage(1, []byte("Zinx hello client"))
		packet := &znet.Packet{}
		data, err := packet.Pack(message)
		if err != nil {
			log.Println("pack message fail:", err)
			return
		}

		if _, err = conn.Write(data); err != nil {
			log.Println("send message fail:", err)
			return
		}

		buffer := make([]byte, packet.HeaderLen())
		if _, err := io.ReadFull(conn, buffer); err != nil {
			log.Println("read message fail:", err)
			return
		}

		message, err = packet.Unpack(buffer)
		if err != nil {
			log.Println("unpack message fail:", err)
			return
		}

		buffer = make([]byte, message.DataLen())
		if _, err := io.ReadFull(conn, buffer); err != nil {
			log.Println("read message fail:", err)
			return
		}

		reader := bytes.NewReader(buffer)
		data = make([]byte, message.DataLen())
		if err := binary.Read(reader, binary.LittleEndian, &data); err != nil {
			log.Println("read message fail:", err)
			return
		}

		log.Printf("server call back: %s [%d]\n", data, message.DataLen())

		time.Sleep(time.Second)
	}
}
