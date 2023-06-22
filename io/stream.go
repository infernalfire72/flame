package io

import (
	"errors"
	"unsafe"
)

type Stream struct {
	Content  []byte
	Position int
}

func NewStream() *Stream {
	return NewStreamWithCapacity(4096)
}

func NewStreamWithCapacity(capacity int) *Stream {
	return &Stream{
		Content: make([]byte, capacity)[:0],
	}
}

func StreamFrom(data []byte) *Stream {
	return &Stream{
		Content: data,
	}
}

func ByteSliceFromPointer(addr unsafe.Pointer, size int) []byte {
	slice := struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{addr, size, size}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

func (s *Stream) WriteAny(p unsafe.Pointer, size int) {
	buffer := ByteSliceFromPointer(p, size)

	if d := len(s.Content) - s.Position; d > 0 {
		for i := 0; i < size && i < d; i++ {
			s.Content[s.Position+i] = buffer[i]
		}

		if len(buffer) > d+1 {
			s.Content = append(s.Content, buffer[d+1:]...)
		}
	} else if d < 0 {
		s.Content = append(s.Content, make([]byte, (-d)+1)...)
		s.Content = append(s.Content, buffer...)
	} else {
		s.Content = append(s.Content, buffer...)
	}

	s.Position += size
}

func (s *Stream) WriteByte(value byte) error {
	if d := len(s.Content) - s.Position; d > 0 {
		s.Content[s.Position] = value
	} else if d < 0 {
		s.Content = append(s.Content, make([]byte, (-d)+1)...)
		s.Content[s.Position] = value
	} else {
		s.Content = append(s.Content, value)
	}

	s.Position++

	return nil
}

func (s *Stream) WriteBoolean(value bool) {
	if value {
		s.WriteByte(1)
	} else {
		s.WriteByte(0)
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
	s.WriteAny(unsafe.Pointer(&value[0]), len(value))
}

func (s *Stream) WriteString(value string) {
	length := len(value)
	if length == 0 {
		s.WriteByte(0)
		return
	}
	s.WriteByte(11)

	for length > 127 {
		s.WriteByte(byte((length & 0x7f) | 0x80))
		length >>= 7
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
	if s.Position+size > len(s.Content) {
		return nil, ErrEndOfStream
	}

	defer func() {
		s.Position += size
	}()

	return unsafe.Pointer(&s.Content[s.Position]), nil
}

func (s *Stream) ReadByte() (byte, error) {
	if s.Position == len(s.Content) {
		return 0, ErrEndOfStream
	}

	defer func() {
		s.Position++
	}()
	return s.Content[s.Position], nil
}

func (s *Stream) ReadSegment(size int) ([]byte, error) {
	if s.Position+size > len(s.Content) {
		return nil, ErrEndOfStream
	}

	result := make([]byte, size)
	if size == 0 {
		return result, nil
	}

	copy(result, s.Content[s.Position:])
	s.Position += size

	return result, nil
}

func (s *Stream) ReadBoolean() (bool, error) {
	if v, err := s.ReadByte(); v >= 1 {
		return true, nil
	} else {
		return false, err
	}
}

func (s *Stream) ReadInt16() (result int16, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(2); err == nil {
		result = *(*int16)(p)
	}
	return
}

func (s *Stream) ReadUint16() (result uint16, err error) {
	var p unsafe.Pointer
	if p, err = s.ReadAny(2); err == nil {
		result = *(*uint16)(p)
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
	if b, _ := s.ReadByte(); b != 0x0b {
		return "", nil
	}

	var (
		length int
		shift  int
		v      byte
		err    error
	)
	for {
		v, err = s.ReadByte()
		if err != nil {
			return "", err
		}
		length |= int(v & 0x7f) << shift
		shift += 7
		if (v & 128) == 0 {
			break
		}
	}

	bytes, err := s.ReadSegment(length)

	return string(bytes), err
}
