package references

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
  "jvm-go/rtda/heap"
)

// invoke a class (static) method
type INVOKE_STATIC struct {
  base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
  cp := frame.Method().Class().ConstantPool()
  methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
  resolveMethod := methodRef.ResolvedMethod()
  if !resolveMethod.IsStatic() {
    panic("java.lang.IncompatibleClassChangeError")
  }
  class := resolveMethod.Class()
  if !class.InitStarted() {
    frame.RevertNextPC()
    base.InitClass(frame.Thread(), class)
    return
  }
  base.InvokeMethod(frame, resolveMethod)
}
