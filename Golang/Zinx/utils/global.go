package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

var GlobalObject = &GlobalObj{}

type GlobalObj struct {
	Server ziface.Server
	Host   string
	Port   int32
	Name   string

	Version        string
	MaxPacketSize  int32
	MaxConn        int32
	WorkerPoolSize int32
	TaskQueueLen   int32
	MaxMsgBufChan  int32

	ConfigPath string
}

func init() {
	GlobalObject = &GlobalObj{
		Host: "0.0.0.0",
		Port: 7777,
		Name: "[Zinx Server]",

		Version:        "v0.0.1",
		MaxPacketSize:  4096,
		MaxConn:        12000,
		WorkerPoolSize: 8,
		TaskQueueLen:   1024,
		MaxMsgBufChan:  64,
	}

	GlobalObject.Reload()
}

func (p *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &GlobalObject); err != nil {
		panic(err)
	}
}
