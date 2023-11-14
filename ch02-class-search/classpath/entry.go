package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
//
// os.PathListSeparator = ';'
const pathListSeparator = string(os.PathListSeparator)

// 表示类路径项的接口
type Entry interface {
	// 负责寻找和加载class文件
	//
	// className- class文件的相对路径
	//
	// 返回读取到的字节数据、最终定位到class文件的Entry、错误消息
	readClass(className string) ([]byte, Entry, error)
	// 相当于java中的toString()
	String() string
}

// 根据路径类型的不同创建Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
