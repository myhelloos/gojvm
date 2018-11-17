package heap

type Object struct {
  class  *Class
  fields Slots
  // todo
}

func (self *Class) NewObject() *Object {
  return newObject(self)
}

func (self *Object) Fields() Slots {
  return self.fields
}
func (self *Object) IsInstanceOf(class *Class) bool {
  return class.isAssignableFrom(self.class)
}

func newObject(class *Class) *Object {
  return &Object{
    class: class,
    fields: newSlots(class.instancesSlotCount),
  }
}