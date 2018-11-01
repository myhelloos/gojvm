package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
  MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute {
  MarkerAttribute
}

type MarkerAttribute {}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
  // return nothing
}
