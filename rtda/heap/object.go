package heap

type Object struct {
  class  *Class
  fields Slots
  // todo
}

func (self *Object) Fields() Slots {
  return self.fields
}

func newObject(class *Class) *Object {
  return &Object{
    class: class,
    fields: newSlots(class.instancesSlotCount),
  }
}