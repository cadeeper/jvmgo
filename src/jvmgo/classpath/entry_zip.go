package classpath

import "errors"
import "archive/zip"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry{
	absDir, err := filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string)([]byte, ZipEntry, error){
	rd, err := zip.OpenReader(self.absDir)
	if err != nil{
		return nil, nil, err
	}
	defer rd.Close()
	for _,file := range rd.File{
		if file.Name == className{
			rc, err := file.Open()
			if err != nil{
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil{
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string{
	return self.absDir
}
