package heap

import (
  "jvm-go/classpath"
  "fmt"
)

type ClassLoader struct {
  cp       *classpath.Classpath
  classMap map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
  return &ClassLoader{
    cp:       cp,
    classMap: make(map[string]*Class),
  }
}

func (self *ClassLoader) LoadClass(name string) *Class {
  if class, ok := self.classMap[name]; ok {
    return class // loaded class
  }
  return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
 data, entry := self.readClass(name)
 class := self.defineClass(data)
 link(class)
 fmt.Printf("[Loaded %s from %s]\n", name, entry)
 return class
}