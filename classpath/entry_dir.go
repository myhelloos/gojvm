package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
  absDir string
}

func newDirEntry(path string) *DirEntry {

}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {

}

func (self *DirEntry) String() string {
  
}
