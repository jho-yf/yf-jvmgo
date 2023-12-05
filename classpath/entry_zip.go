package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

// ZIP或JAR文件形式的类路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 遍历压缩包里面的文件，查找class文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
