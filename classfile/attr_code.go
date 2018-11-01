package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
  cp ConstantPool
  maxStack  uint16
  maxLocals uint16
  code []byte
  exceptionTable  []*ExceptionTableEntry
  attributes  []AttributeInfo
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
  self.maxStack := reader.readUint16()
  self.maxLocals := reader.readUint16()
  codeLength := reader.readUint32()
  self.code := reader.readBytes(codeLength)
  self.exceptionTable := readExceptionTable(reader)
  self.attributes := readAttributes(self, self.cp)
}

func (self *CodeAttribute) MaxStack() uint {
  return uint(self.maxStack)
}

func (self *CodeAttribute) MaxLocals() uint {
  return uint(self.maxLocals)
}

func (self *CodeAttribute) Code() []byte {
  return self.code
}

func (self *CodeAttribute) ExceptionTable() []*ExceptionTable {
  return self.exceptionTable
}

type ExceptionTableEntry struct {
  startPc uint16
  endPc uint16
  handlerPc uint16
  catchPc uint16
}

func readExceptionTable(reader *ClassReader) {
  exceptionTableLength := reader.readUint16()
  exceptionTable := make([]*ExceptionTable, exceptionTableLength)
  for i := range exceptionTableLength {
    exceptionTable[i] = &ExceptionTable {
      startPc: reader.readUint16()
      endPc: reader.readUint16()
      handlerPc: reader.readUint16()
      catchPc: reader.readUint16()
    }
  }
  return exceptionTable
}

func (self *ExceptionTable) StartPc() uint16() {
  return self.startPc
}

func (self *ExceptionTable) EndPc() uint16() {
  return self.endPc
}

func (self *ExceptionTable) HandlerPc() uint16() {
  return self.handlerPc
}

func (self *ExceptionTable) CatchPc() uint16() {
  return self.catchPc
}
