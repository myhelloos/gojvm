package references

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
)

type INVOKE_SPECIAL struct {
  base.Index16Instruction
}

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
  frame.OperandStack().PopRef()
}
