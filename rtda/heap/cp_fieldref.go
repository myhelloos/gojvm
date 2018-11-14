package heap

import "jvm-go/classfile"

type FieldRef struct  {
  MemberRef
  field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
  ref := &FieldRef{}
  ref.cp = cp
  ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
  return ref
}

func (self *FieldRef) ResolvedField() *Field {
  if self.field == nil {
    self.resolveClassRef()
  }
  return self.field
}

func (self *FieldRef) resolveClassRef() {
  d := self.cp.class
  c := self.ResolvedClass()
  field := lookupFiled(c, self.name, self.descriptor)

  if field == nil {
    panic("java.lang.NoSuchFieldError")
  }
  if !field.isAccessibleTo(d) {
    panic("java.lang.IllegalAccessError")
  }

  self.field = field
}
func lookupFiled(c *Class, name, descriptor string) *Field {
  for _, field := range c.fields {
    if field.name == name && field.descriptor == descriptor {
      return field
    }
  }
  for _, iface := range c.interfaces {
    if field := lookupFiled(iface, name, descriptor); field != nil {
      return field
    }
  }
  if c.superClass != nil {
    return lookupFiled(c.superClass, name, descriptor)
  }
  return nil
}