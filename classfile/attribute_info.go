package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
  readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
  attributeCount := reader.readUint16()
  attributes := make([]AttributeInfo, attributeCount)
  for i := range attributes {
    attributes[i] = readAttribute(reader, cp)
  }
  return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
  attributeNameIndex := reader.readUint16()
  attributeName := cp.getUtf8(attributeNameIndex)
  attributeLength := reader.readUint32()
  attributeInfo := newAttributeInfo(attributeName, attributeLength, cp)
  attributeInfo.readInfo(reader)
  return attributeInfo
}

func newAttributeInfo(attributeName string, attributeLength uint32, cp ConstantPool) AttributeInfo {
  switch attributeName {
  case "Code":
    return &CodeAttribute{cp: cp}
  case "ConstantValue":
    return &ConstantValueAttribute{}
  case "Deprecated":
    return &DeprecatedAttribute{}
  case "Exceptions":
    return &ExceptionsAttribute{}
  case "LineNumberTable":
    return &LineNumberTableAttribute{}
  case "LocalVariableTable":
    return &LocalVariableTableAttribute{}
  case "SourceFile":
    return &SourceFileAttribute{cp: cp}
  case "Synthetic":
    return &SyntheticAttribute{}
  default:
    return &UnparsedAttribute{attributeName, attributeLength, nil}
  }
}
