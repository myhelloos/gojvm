package lang

import (
  "jvm-go/native"
  "jvm-go/rtda"
  "jvm-go/rtda/heap"
)

func init() {
  native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int desPost, int length)
func arraycopy(frame *rtda.Frame) {
  vars := frame.LocalVars()
  src := vars.GetRef(0)
  srcPos := vars.GetInt(1)
  dst := vars.GetRef(2)
  dstPos := vars.GetInt(3)
  length := vars.GetInt(4)

  if src == nil || dst == nil {
    panic("java.lang.NullPointerException")
  }

  if !checkArrayCopy(src, dst) {
    panic("java.lang.ArrayStoreException")
  }

  if srcPos < 0 || dstPos < 0 || length < 0 ||
    srcPos+length > src.ArrayLength() ||
    dstPos+length > dst.ArrayLength() {
    panic("java.lang.IndexOutOfBoundException")
  }

  heap.ArrayCopy(src, dst, srcPos, dstPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
  srcClass := src.Class()
  destClass := dest.Class()
  if !srcClass.IsArray() || !destClass.IsArray() {
    return false
  }
  if srcClass.ComponentClass().IsPrimitive() ||
    destClass.ComponentClass().IsPrimitive() {
    return srcClass == destClass
  }
  return true
}
