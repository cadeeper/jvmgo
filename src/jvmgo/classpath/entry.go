package classpath
import "os"
import "strings"

const pathListSeperator = string(os.PathListSeparator)

type Entry interface {
	//path must be like path/to/object.class
	readClass(className string)([]byte, Entry, error)
	String() string
}


func newEntry(path string) Entry{
	if strings.Contains(path, pathListSeperator){
		return newCompositeEntry(path)
	}else if strings.HasSuffix(path, "*"){
		return newWildcardEntry(path)
	}else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR"){
		return newZipEntry(path)
	}else if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

