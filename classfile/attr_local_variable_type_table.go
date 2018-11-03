package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTable struct {
  localVariableTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
  startPc uint16
  length uint16
  nameIndex uint16
  signatureIndex uint16
  index uint16
}

func (self *LocalVariableTypeTable) readInfo(reader *ClassReader) {
  localVariableTableLength := reader.readUint16()
  self.localVariableTable = make([]*LocalVariableTypeTableEntry, localVariableTableLength)
  for i := range self.localVariableTable {
    self.localVariableTable[i] = &LocalVariableTypeTableEntry {
      startPc: reader.readUint16(),
      length: reader.readUint16(),
      nameIndex: reader.readUint16(),
      signatureIndex: reader.readUint16(),
      index: reader.readUint16(),
    }
  }
}
