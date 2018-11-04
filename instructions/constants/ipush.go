package constants

import (
  "jvm-go/rtda"
  "jvm-go/instructions/base"
)

// push byte
type BIPUSH struct {
  val int8
}
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
  self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
  i := int32(self.val)
  frame.OperandStack().PushInt(i)
}

// push short
type SIPUSH struct {
  val int16
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
  self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
  i := int32(self.val)
  frame.OperandStack().PushInt(i)
}
