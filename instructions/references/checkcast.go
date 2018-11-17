package references

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
  "jvm-go/rtda/heap"
)

// check whether object is of given type
type CHECK_CAST struct {
  base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  ref := stack.PopRef()
  stack.PushRef(ref)

  if ref == nil {
    return
  }

  cp := frame.Method().Class().ConstantPool()
  classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
  class := classRef.ResolvedClass()

  if !ref.IsInstanceOf(class) {
    panic("java.lang.ClassCastException")
  }
}
