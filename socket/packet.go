package socket

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const headLen = 8

type handler struct{}

func (h *handler) UnpackData(bin []byte) (cmd uint32, message []byte, err error) {
	if len(bin) < headLen {
		err = errors.New("data size err")
		return
	}
	cmd = binary.BigEndian.Uint32(bin[0:])
	message = bin[3:]
	err = nil
	return
}

func (h *handler) PackData(code uint32, cmd uint32, data []byte) []byte {
	buffer := bytes.NewBuffer([]byte{})
	size := uint16(len(data) + headLen)
	binary.Write(buffer, binary.BigEndian, size)
	binary.Write(buffer, binary.BigEndian, code)
	binary.Write(buffer, binary.BigEndian, cmd)
	binary.Write(buffer, binary.BigEndian, data)
	return buffer.Bytes()
}
