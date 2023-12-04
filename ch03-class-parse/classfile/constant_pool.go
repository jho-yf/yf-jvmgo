package classfile

// 常量池占据class文件很大一部分数据，里面存放着各式各样的常量信息
// 包括：数字、字符串常量、类和接口名等
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return nil
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	return "", ""
}

func (cp ConstantPool) getClassName(index uint16) string {
	return ""
}

func (cp ConstantPool) getUtf8(index uint16) string {
	return ""
}
