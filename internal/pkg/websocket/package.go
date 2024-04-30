package websocket

import (
	"bytes"
	"encoding/binary"
)

const CMessageLenBit = 4

type MessagePack struct {
	bufPool    []byte
	isContinue bool
}

func MarshalPack(src []byte) []byte {
	length := len(src)
	buf := bytes.NewBuffer(make([]byte, 0, length+4))
	binary.Write(buf, binary.LittleEndian, int32(length))
	buf.Write(src)
	return buf.Bytes()
}

func (p *MessagePack) UnmarshalPack(src []byte) ([]byte, bool) {
	p.bufPool = append(p.bufPool, src...)
	if len(p.bufPool) < CMessageLenBit {
		return nil, false
	}

	messageLen := int(binary.LittleEndian.Uint32(src))

	if len(p.bufPool) < messageLen+CMessageLenBit {
		return nil, false
	}

	messageBuf := p.bufPool[CMessageLenBit : messageLen+CMessageLenBit]
	p.bufPool = p.bufPool[messageLen+CMessageLenBit:]
	if len(p.bufPool) >= CMessageLenBit {
		return messageBuf, true
	}

	return messageBuf, false

}
