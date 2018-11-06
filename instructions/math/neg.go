package math

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
)

// Negate double
type DNEG struct {
  base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v1 := stack.PopDouble()
  stack.PushDouble(-v1)
}

// Negate float
type FNEG struct {
  base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v1 := stack.PopFloat()
  stack.PushFloat(-v1)
}

// Negate int
type INEG struct {
  base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v1 := stack.PopInt()
  stack.PushInt(-v1)
}

// Negate long
type LNEG struct {
  base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v1 := stack.PopLong()
  stack.PushLong(-v1)
}