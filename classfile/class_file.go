package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
  magic         uint32
  minorVersion  uint16
  majorVersion  uint16
  constantPool  ConstantPool
  accessFlags   uint16
  thisClass     uint16
  superClass    uint16
  interfaces    []uint16
  fields        []*MemberInfo
  methods       []*MemberInfo
  attributes    []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
  defer func() {
    if r := recover(); r != nil {
      var ok bool
      err, ok = r.(error)
      if !ok {
        err = fmt.Errorf("%v", r)
      }
    }
  }()

  cr := &ClassReader(classData)
  cf = &ClassFile{}
  cf.read(cr)
  return
}

func (self *ClassFile) read(reader *ClassReader) {
  self.readAndCheckMagic(reader)
  self.readAndCheckVersion(version)
  self.constantPool = readConstantPool(reader)
  self.accessFlags = reader.readUint16()
  self.thisClass = reader.readUint16()
  self.superClass = reader.readUint16()
  self.interfaces = reader.readUint16s()
  self.fields = readMembers(reader, self.constantPool)
  self.methods = readMembers(reader, self.constantPool)
  self.attributes = readAttributes(reader, self.constantPool)
}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {}
// getter
func (self *ClassFile) MinorVersion() uint16{
  return self.minorVersion
}
// getter
func (self *ClassFile) MajorVersion() uint16{
  return self.majorVersion
}
// getter
func (self *ClassFile) ConstantPool() ConstantPool{
  return self.constantPool
}
// getter
func (self *ClassFile) AccessFlags() uint16{
  return self.accessFlags
}
// getter
func (self *ClassFile) Fields() *MemberInfo {
  return self.fields
}
// getter
func (self *ClassFile) Methods() *MemberInfo{
  return self.methods
}
// getter
func (self *ClassFile) ClassName() string {
  return self.constantPool.getClassName(self.thisClass)
}
// getter
func (self *ClassFile) SuperClassName() string {
  if self.superClass > 0 {
    return self.constantPool.getClassName(self.superClass)
  }
  return "" // java.lang.Object
}
// getter
func (self *ClassFile) InterfaceName() []string {
  interfaceNames := make([]string, len(self.interfaces))
  for i, cpIndex := range self.interfaces {
    interfaceNames[i] = self.constantPool.getClassName(cpIndex)
  }
  return interfaceNames
}
