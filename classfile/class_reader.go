package classfile

import "encoding/binary"

type ClassReader struct {
  data []byte
}

// read u1
func (self *ClassReader) readUnit8() uint8 {

}

// read u2
func (self *ClassReader) readUnit16() uint16 {

}

// read u4
func (self *ClassReader) readUnit32() uint32 {

}

func (self *ClassReader) readUnit64() uint64 {

}

func (self *ClassReader) readUnit16s() []uint16 {

}

func (self *ClassReader) readBytes(length uint32) []byte {

}
