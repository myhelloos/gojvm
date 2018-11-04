package constants

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
)

// do nothing
type NOP struct {
  base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
  // really do nothing
}