package references

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
  "jvm-go/rtda/heap"
)

type INVOKE_SPECIAL struct {
  base.Index16Instruction
}

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
  currentClass := frame.Method().Class()
  cp := currentClass.ConstantPool()
  methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
  resolvedClass := methodRef.ResolvedClass()
  resolvedMehtod := methodRef.ResolvedMethod()

  if resolvedMehtod.Name() == "<init>" && resolvedMehtod.Class() != resolvedClass {
    panic("java.lang.NoSuchMethod")
  }
  if resolvedMehtod.IsStatic() {
    panic("java.lang.IncompatibleClassChangeError")
  }

  ref := frame.OperandStack().GetRefFromTop(resolvedMehtod.ArgSlotCount() - 1)
  if ref == nil {
    panic("java.lang.NullPointerException")
  }

  if resolvedMehtod.IsProtected() &&
    resolvedMehtod.Class().IsSuperClassOf(currentClass) &&
    resolvedMehtod.Class().GetPackageName() != currentClass.GetPackageName() &&
    ref.Class() != currentClass &&
    !ref.Class().IsSubClassOf(currentClass) {
    panic("java.lang.IllegalAccessError")
  }

  methodToBeInvoked := resolvedMehtod
  if currentClass.IsSuper() &&
    resolvedClass.IsSuperClassOf(currentClass) &&
    resolvedMehtod.Name() != "<init>" {
    methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
  }

  if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
    panic("java.lang.AbstractMethodRef")
  }

  base.InvokeMethod(frame, methodToBeInvoked)
}
