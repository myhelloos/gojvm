package heap

import "jvm-go/classfile"

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
  accessFlags uint16
  name string // thisClassName
  superClassName string
  interfaceNames []string
  constantPool *ConstantPool
  fields []*Field
  methods []*Method
  loader *ClassLoader
  superClass *Class
  interfaces []*Class
  instancesSlotCount uint
  staticSlotCount uint
  staticVars Slots
}

func newClass(cf *classfile.ClassFile) *Class {
  class := &Class{}
  class.accessFlags = cf.AccessFlags()
  class.name = cf.ClassName()
  class.superClassName = cf.SuperClassName()
  class.interfaceNames = cf.InterfaceNames()
  class.constantPool = newConstantPool(class, cf.ConstantPool())
  class.fields = newFields(class, cf.Fields())
  class.methods = newMethods(class, cf.Methods())
  return class
}

func (self *Class) IsPublic() bool {
  return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
  return 0 != self.accessFlags & ACC_FINAL
}
func (self *Class) IsSuper() bool {
  return 0 != self.accessFlags & ACC_SUPER
}
func (self *Class) IsInterface() bool {
  return 0 != self.accessFlags & ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
  return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
  return 0 != self.accessFlags & ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
  return 0 != self.accessFlags & ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
  return 0 != self.accessFlags & ACC_ENUM
}

// getter
func (self *Class) ConstantPool() *ConstantPool {
  return self.constantPool
}

func (self *Class) StaticVars() Slots {
  return self.staticVars
}