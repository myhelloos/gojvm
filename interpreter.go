package main

import (
  "jvm-go/classfile"
  "jvm-go/rtda"
  "fmt"
  "jvm-go/instructions/base"
  "jvm-go/instructions"
)

func interpret(methodInfo *classfile.MemberInfo) {
  codeAttr := methodInfo.CodeAttribute()
  maxLocals := codeAttr.MaxLocals()
  maxStack := codeAttr.MaxStack()
  byteCode := codeAttr.Code()

  thread := rtda.NewThread()
  frame := thread.NewFrame(maxLocals, maxStack)
  thread.PushFrame(frame)

  defer catchErr(frame)
  loop(thread, byteCode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
  frame := thread.PopFrame()
  reader := &base.BytecodeReader{}
  for {
    pc := frame.NextPC()
    thread.SetPC(pc)

    // decode
    reader.Reset(bytecode, pc)
    opcode := reader.ReadUint8()
    inst := instructions.NewInstruction(opcode)
    inst.FetchOperands(reader)
    frame.SetNextPC(reader.PC())

    // execute
    fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
    inst.Execute(frame)
  }
}

func catchErr(frame *rtda.Frame) {
  if r := recover(); r != nil {
    fmt.Printf("LocalVars: %V\n", frame.LocalVars())
    fmt.Printf("OperandStack: %V\n", frame.OperandStack())
    panic(r)
  }
}