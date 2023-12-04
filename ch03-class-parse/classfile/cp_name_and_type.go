package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (nameAndTypeInfo *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	nameAndTypeInfo.nameIndex = reader.readUint16()
	nameAndTypeInfo.descriptorIndex = reader.readUint16()
}
