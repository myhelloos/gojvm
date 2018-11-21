package heap

type Object struct {
  class *Class
  data  interface{}
  extra interface{}
}

func newObject(class *Class) *Object {
  return &Object{
    class: class,
    data:  newSlots(class.instancesSlotCount),
  }
}

func (self *Object) Fields() Slots {
  return self.data.(Slots)
}
func (self *Object) Class() *Class {
  return self.class
}
func (self *Object) Extra() interface{} {
  return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
  self.extra = extra
}
func (self *Object) IsInstanceOf(class *Class) bool {
  return class.isAssignableFrom(self.class)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
  filed := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  slots.SetRef(filed.slotId, ref)
}
func (self *Object) GetRefVar(name, descriptor string) *Object {
  filed := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  return slots.GetRef(filed.slotId)
}
