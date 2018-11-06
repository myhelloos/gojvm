package math

import (
  "jvm-go/rtda"
  "jvm-go/instructions/base"
)

// Increment local variable by constant
type IINC struct {
  Index uint
  Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
  self.Index = uint(reader.ReadUint8())
  self.Const = reader.ReadInt32()
}

func (self *IINC) Execute(frame *rtda.Frame) {
  localVars := frame.LocalVars()
  val := localVars.GetInt(self.Index)
  val += self.Const
  localVars.SetInt(self.Index, val)
}