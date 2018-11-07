package control

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
)

type LOOKUP_SWITCH struct {
  defaultOffset int32
  npair int32
  matchOffsets []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
  reader.SkipPadding()
  self.defaultOffset = reader.ReadInt32()
  self.npair = reader.ReadInt32()
  self.matchOffsets = reader.ReadInt32s(self.npair * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
  key := frame.OperandStack().PopInt()
  for i := int32(0); i < self.npair * 2; i += 2 {
    if self.matchOffsets[i] == key {
      offset := self.matchOffsets[i+1]
      base.Branch(frame, int(offset))
      return
    }
  }
  base.Branch(frame, int(self.defaultOffset))
}