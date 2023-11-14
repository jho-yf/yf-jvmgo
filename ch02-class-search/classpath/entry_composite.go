package classpath

import (
	"errors"
	"strings"
)

// 表示组合形式的类路径
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	// 将路径列表参数按照分隔符分割
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历每个路径，找到class文件并读取，如果读取不到返回错误
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strArr := make([]string, len(self))
	for i, entry := range self {
		strArr[i] = entry.String()
	}
	return strings.Join(strArr, pathListSeparator)
}
