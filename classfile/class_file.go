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

func Parse(classData []byte) (cf *ClassFile, err error) {}

func (self *ClassFile) read(reader *ClassReader) {}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {}
// getter
func (self *ClassFile) MinorVersion() uint16{}
// getter
func (self *ClassFile) MajorVersion() uint16{}
// getter
func (self *ClassFile) ConstantPool() ConstantPool{}
// getter
func (self *ClassFile) AccessFlags() uint16{}
// getter
func (self *ClassFile) Fields() *MemberInfo {}
// getter
func (self *ClassFile) Methods() *MemberInfo{}
// getter
func (self *ClassFile) ClassName() string {

}
func (self *ClassFile) SuperClassName() string {

}
func (self *ClassFile) InterfaceName() []string {

}
