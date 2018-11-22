package control

import (
  "jvm-go/rtda"
  "jvm-go/instructions/base"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
  defaultOffset int32
  low           int32
  high          int32
  jumpOffset    []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
  reader.SkipPadding()
  self.defaultOffset = reader.ReadInt32()
  self.low = reader.ReadInt32()
  self.high = reader.ReadInt32()
  jumpOffsetCount := self.high - self.low + 1
  self.jumpOffset = reader.ReadInt32s(jumpOffsetCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
  index := frame.OperandStack().PopInt()

  var offset int
  if index >= self.low && index <= self.high {
    offset = int(self.jumpOffset[index-self.low])
  } else {
    offset = int(self.defaultOffset)
  }

  base.Branch(frame, offset)
}
