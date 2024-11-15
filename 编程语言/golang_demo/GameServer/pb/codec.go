package pb

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	sync "sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

var (
	codeToDescriptor map[int32]protoreflect.MessageDescriptor
	nameToCode       map[string]int32
	once             sync.Once
)

func Init() {
	once.Do(func() {
		codeToDescriptor = make(map[int32]protoreflect.MessageDescriptor)
		nameToCode = make(map[string]int32)

		for code, name := range MsgCode_name {
			name = strings.ToLower(strings.ReplaceAll(name, "_", ""))
			nameToCode[name] = code
		}

		descriptors := File_proto_game_proto.Messages()
		for i := 0; i < descriptors.Len(); i++ {
			descriptor := descriptors.Get(i)
			name := strings.ToLower(strings.ReplaceAll(string(descriptor.Name()), "_", ""))
			if code, ok := nameToCode[name]; ok {
				codeToDescriptor[code] = descriptor
			}
		}

	})
}

func Decode(data []byte, code int32) (proto.Message, error) {
	if data == nil {
		return nil, errors.New("data is empty")
	}

	descriptor, err := getMessageDescriptor(code)
	if err != nil {
		return nil, err
	}

	message := dynamicpb.NewMessage(descriptor)
	if err := proto.Unmarshal(data, message); err != nil {
		return nil, err
	}

	return message, nil
}

func Encode(message proto.Message) ([]byte, error) {
	if message == nil {
		return nil, errors.New("message is nil")
	}

	code, err := getMessageCode(message)
	if err != nil {
		return nil, err
	}

	codeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(codeBytes, uint16(code))

	msgBytes, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(len(msgBytes)))

	data := append(lengthBytes, codeBytes...)
	data = append(data, msgBytes...)

	return data, nil
}

func getMessageDescriptor(code int32) (protoreflect.MessageDescriptor, error) {
	descriptor, ok := codeToDescriptor[code]
	if !ok {
		return nil, fmt.Errorf("unknown code %d", code)
	}

	return descriptor, nil
}

func getMessageCode(message proto.Message) (int32, error) {
	if message == nil {
		return 0, errors.New("message is nil")
	}

	var (
		descriptor = message.ProtoReflect().Descriptor()
		name       = strings.ToLower(strings.ReplaceAll(string(descriptor.Name()), "_", ""))
		code, ok   = nameToCode[name]
	)

	if !ok {
		return 0, fmt.Errorf("unknown message %s", descriptor.Name())
	}

	return code, nil
}
