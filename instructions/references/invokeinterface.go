package references

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
  "jvm-go/rtda/heap"
)

// invoke interface method
type INVOKE_INTERFACE struct {
  index uint
  // count uint8
  // zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
  self.index = uint(reader.ReadUint16())
  reader.ReadUint8() // count
  reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
  cp := frame.Method().Class().ConstantPool()
  methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
  resolveMethod := methodRef.ResolvedInterfaceMethod()
  if resolveMethod.IsStatic() || resolveMethod.IsPrivate() {
    panic("java.lang.IncompatibleClassChangeError")
  }
  ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount() - 1)
  if ref == nil {
    panic("java.lang.NullPointerException")
  }
  if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
    panic("java.lang.IncompatibleClassChangeError")
  }

  methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
  if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
    panic("java.lang.AbstractMethodError")
  }
  if !methodToBeInvoked.IsPublic() {
    panic("java.lang.IllegalAccessError")
  }

  base.InvokeMethod(frame, methodToBeInvoked)
}
