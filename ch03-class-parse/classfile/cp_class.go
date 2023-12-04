package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (classInfo *ConstantClassInfo) readInfo(reader *ClassReader) {
	classInfo.nameIndex = reader.readUint16()
}
