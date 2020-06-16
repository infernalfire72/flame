package io

import (
	"errors"
	"unsafe"
)

type Stream struct {
	Content		[]byte
	Length		int
	Capacity	int
	Position	int
}

func NewStream() *Stream {
	return NewStreamWithCapacity(4096)
}

func NewStreamWithCapacity(capacity int) *Stream {
	return &Stream {
		Content:	make([]byte, capacity),
		Capacity:	capacity,
	}
}

func StreamFrom(data []byte) *Stream {
	return &Stream {
		Content:	data,
		Capacity:	len(data),
		Length:		len(data),
	}
}

func (s *Stream) Data() []byte {
	return s.Content[:s.Length]
}

func (s *Stream) Realloc(capacity int) {
	buffer := make([]byte, capacity)
	copy(buffer, s.Content)
	s.Content = buffer
	s.Capacity = capacity
}

func ByteSliceFromPointer(addr unsafe.Pointer, size int) []byte {
	slice := struct {
		addr	unsafe.Pointer
		len		int
		cap		int
	}{addr, size, size}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

func (s *Stream) WriteAny(p unsafe.Pointer, size int) {
	if s.Length + size > s.Capacity {
		s.Realloc(s.Capacity * 2)
	}

	buffer := ByteSliceFromPointer(p, size)
	copy(s.Content[s.Position:], buffer)

	s.Position += size
	if s.Position > s.Length {
		s.Length = s.Position
	}
}

func (s *Stream) WriteByte(value byte) {
	if s.Length + 1 > s.Capacity {
		s.Realloc(s.Capacity * 2)
	}

	s.Content[s.Position] = value

	s.Position++
	if s.Position > s.Length {
		s.Length = s.Position
	}
}

func (s *Stream) WriteInt16(value int16) {
	s.WriteAny(unsafe.Pointer(&value), 2)
}

func (s *Stream) WriteInt32(value int32) {
	s.WriteAny(unsafe.Pointer(&value), 4)
}

func (s *Stream) WriteInt64(value int64) {
	s.WriteAny(unsafe.Pointer(&value), 8)
}

func (s *Stream) WriteFloat32(value float32) {
	s.WriteAny(unsafe.Pointer(&value), 4)
}

func (s *Stream) WriteFloat64(value float64) {
	s.WriteAny(unsafe.Pointer(&value), 8)
}

func (s *Stream) WriteByteSlice(value []byte) {
	if s.Length + len(value) > s.Capacity {
		s.Realloc(s.Capacity * 2)
	}

	copy(s.Content[s.Position:], value)

	s.Position += len(value)
	if s.Position > s.Length {
		s.Length = s.Position
	}
}

func (s *Stream) WriteString(value string) {
	length := len(value)
	if length == 0 {
		s.WriteByte(0)
		return
	}
	s.WriteByte(11)

	for length >= 127 {
		s.WriteByte(128)
		length -= 127
	}
	s.WriteByte(byte(length))
	s.WriteByteSlice([]byte(value))
}

func (s *Stream) WriteInterface(value interface{}) {
	switch v := value.(type) {
	case int32:
		s.WriteInt32(v)
	case int:
		s.WriteInt32(int32(v))
	case int16:
		s.WriteInt16(v)
	case int64:
		s.WriteInt64(v)
	case float32:
		s.WriteFloat32(v)
	case float64:
		s.WriteFloat64(v)
	case byte:
		s.WriteByte(v)
	case []byte:
		s.WriteByteSlice(v)
	case string:
		s.WriteString(v)
	}
}

var ErrEndOfStream error = errors.New("io: end of stream reached")

func (s *Stream) ReadAny(size int) (unsafe.Pointer, error) {
	if s.Position + size > s.Length {
		return nil, ErrEndOfStream
	}

	defer func() {
		s.Position += size
	}()

	return unsafe.Pointer(&s.Content[s.Position]), nil
}

func (s *Stream) ReadByte() int {
	if s.Position == s.Length {
		return -1
	}

	defer func() {
		s.Position++
	}()
	return int(s.Content[s.Position])
}

func (s *Stream) ReadSegment(size int) ([]byte, error) {
	if s.Position + size > s.Length {
		return nil, ErrEndOfStream
	}

	result := make([]byte, size)
	if size == 0 {
		return result, nil
	}

	copy(result, s.Content[s.Position:]) // We don't need a second delimiter because it only copies for dst size
	s.Position += size

	return result, nil
}

func (s *Stream) ReadInt16() (result int16, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(2); err == nil {
		result = *(*int16)(p)
	}
	return
}

func (s *Stream) ReadInt32() (result int32, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(4); err == nil {
		result = *(*int32)(p)
	}
	return
}

func (s *Stream) ReadInt64() (result int64, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(8); err == nil {
		result = *(*int64)(p)
	}
	return
}

func (s *Stream) ReadFloat32() (result float32, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(4); err == nil {
		result = *(*float32)(p)
	}
	return
}

func (s *Stream) ReadFloat64() (result float64, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(8); err == nil {
		result = *(*float64)(p)
	}
	return
}

func (s *Stream) ReadString() (string, error) {
	if s.ReadByte() != 0x0b {
		return "", nil
	}

	length := 0
	var v int
	for {
		v = s.ReadByte()
		if v == 128 {
			length += 127
		} else if v == -1 {
			return "", ErrEndOfStream
		} else {
			length += v
			break
		}
	}

	bytes, err := s.ReadSegment(length)

	return string(bytes), err
}