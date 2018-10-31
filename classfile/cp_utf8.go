package classfile

import "fmt"
import "unicode/utf16"

type ConstantUf8Info struct {
  str string
}

func (self *ConstantUf8Info) readInfo(reader *ClassReader) {
  length := uint32(reader.readUint16())
  bytes := reader.readBytes(length)
  self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
  return string(bytes)
}
