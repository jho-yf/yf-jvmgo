package classpath

import (
	"os"
	"path/filepath"
)

// 表示目录形式的类路径
type DirEntry struct {
	// 目录的绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 将路径转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 拼接目录和文件名成为一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 读取class文件内容
	data, err := os.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
