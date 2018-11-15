package rtda

import "jvm-go/rtda/heap"

type Frame struct {
  lower *Frame
  localVars LocalVars
  operandStack *OperandStack
  thread *Thread
  method *heap.Method
  nextPC int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
  return &Frame{
    thread: thread,
    localVars: newLocalVars(maxLocals),
    operandStack: newOperandStack(maxStack),
  }
}

// getters
func (self *Frame) LocalVars() LocalVars {
  return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
  return self.operandStack
}
func (self *Frame) NextPC() int {
  return self.nextPC
}
func (self *Frame) Thread() *Thread {
  return self.thread
}
func (self *Frame) Method() *heap.Method {
  return self.method
}

// setter
func (self *Frame) SetNextPC(nextPC int) {
  self.nextPC = nextPC
}

