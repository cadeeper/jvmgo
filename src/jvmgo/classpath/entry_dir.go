package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry{
	absDir, err := filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string)([]byte, DirEntry, error){
	filename := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string{
	return self.absDir
}
