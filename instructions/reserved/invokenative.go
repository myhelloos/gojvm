package reserved

import (
  "jvm-go/instructions/base"
  "jvm-go/rtda"
  "jvm-go/native"
  _ "jvm-go/native/java/lang"
  _ "jvm-go/native/java/sun/misc"
)

type INVOKE_NATIVE struct {
  base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
  method := frame.Method()
  className := method.Class().Name()
  methodName := method.Name()
  methodDescriptor := method.Descriptor()

  nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
  if nativeMethod == nil {
    methodInfo := className + "." + methodName + methodDescriptor
    panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
  }
  nativeMethod(frame)
}
