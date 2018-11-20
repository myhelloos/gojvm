package heap

type Object struct {
  class *Class
  data  interface{}
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
