package control

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
)

// return void from method
type RETURN struct {
  base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
  frame.Thread().PopFrame()
}

// return reference from method
type ARETURN struct {
  base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
  thread := frame.Thread()
  currentFrame := thread.PopFrame()
  invokerFrame := thread.TopFrame()
  retVal := currentFrame.OperandStack().PopRef()
  invokerFrame.OperandStack().PushRef(retVal)
}

// return double from method
type DRETURN struct {
  base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
  thread := frame.Thread()
  currentFrame := thread.PopFrame()
  invokerFrame := thread.TopFrame()
  retVal := currentFrame.OperandStack().PopDouble()
  invokerFrame.OperandStack().PushDouble(retVal)
}

// return float from method
type FRETURN struct {
  base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
  thread := frame.Thread()
  currentFrame := thread.PopFrame()
  invokerFrame := thread.TopFrame()
  retVal := currentFrame.OperandStack().PopFloat()
  invokerFrame.OperandStack().PushFloat(retVal)
}

// return int from method
type IRETURN struct {
  base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
  thread := frame.Thread()
  currentFrame := thread.PopFrame()
  invokerFrame := thread.TopFrame()
  retVal := currentFrame.OperandStack().PopInt()
  invokerFrame.OperandStack().PushInt(retVal)
}

// return long from method
type LRETURN struct {
  base.NoOperandsInstruction
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
  thread := frame.Thread()
  currentFrame := thread.PopFrame()
  invokerFrame := thread.TopFrame()
  retVal := currentFrame.OperandStack().PopLong()
  invokerFrame.OperandStack().PushLong(retVal)
}
