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

	}else if strings.HasSuffix(path, "*"){

	}else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR"){

	}else if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"){

	}
	return
}

