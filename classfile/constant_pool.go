package classfile

// 常量池占据class文件很大一部分数据，里面存放着各式各样的常量信息
// 包括：数字、字符串常量、类和接口名等
type ConstantPool []ConstantInfo

// 常量池也是一个表，表头给出的值是常量池大小，假设表头给出的值是n
//
// 1. 常量池表头给出的常量池大小比实际大1，有效的常量池索引是1到n-1，0为无效索引,不指向任何常量
// 2. CONSTANT_Long_info 和 CONSTANT_Double_info各占两个位置，所以常量数量可能比n-1还少
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	// 索引从1开始，0为无效索引,不指向任何常量
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// CONSTANT_Long_info 和 CONSTANT_Double_info各占两个位置
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// 按照索引查找常量
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

// 从常量池查找字段或方法的字符和描述符
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池中查找类名
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

// 从常量池中查找UTF-8字符串
func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
