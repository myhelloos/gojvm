package classfile

import "encoding/binary"

type ClassReader struct {
  data []byte
}

// read u1
func (self *ClassReader) readUint8() uint8 {
  val := self.data[0]
  self.data = self.data[1:]
  return val
}

// read u2
func (self *ClassReader) readUint16() uint16 {
  val := binary.BigEndian.Unit16(self.data)
  self.data = self[2:]
  return val
}

// read u4
func (self *ClassReader) readUint32() uint32 {
  val := binary.BigEndian.Uint32(self.data)
  self.data = self[4:]
  return val
}

func (self *ClassReader) readUint64() uint64 {
  val := binary.BigEndian.Uint64(self.data)
  self.data = self[8:]
  return val
}

func (self *ClassReader) readUint16s() []uint16 {
  n := self.readUint16()
  s := make([]unit16, n)
  for i := range s {
    s[i] = self.readUint16()
  }
  return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
  bytes := self.data[:n]
  self.data = self.data[n:]
  return bytes
}
