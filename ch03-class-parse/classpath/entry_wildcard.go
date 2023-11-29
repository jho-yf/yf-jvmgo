package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去掉末尾的星号，得到baseDir
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过子目录遍历（通配符类路径不能递归匹配子目录下的JAR文件）
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 根据后缀名选出JAR文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	// 遍历baseDir，创建ZipEntry
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
