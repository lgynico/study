package ziface

type Message interface {
	ID() int32
	DataLen() int32
	Data() []byte

	SetID(id int32)
	SetDataLen(len int32)
	SetData(data []byte)
}
