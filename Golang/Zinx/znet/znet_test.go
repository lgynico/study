package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	msg := NewMessage(1, []byte("Hello World"))
	pack := &Packet{}
	data, err := pack.Pack(msg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)

	msg, err = pack.Unpack(data)
	if err != nil {
		t.Fatal(err)
	}

	reader := bytes.NewReader(data[pack.HeaderLen():])
	data = make([]byte, msg.DataLen())
	err = binary.Read(reader, binary.LittleEndian, &data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}
