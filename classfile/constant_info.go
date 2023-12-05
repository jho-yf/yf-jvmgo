package classfile

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(cr *ClassReader, cp ConstantPool) ConstantInfo {
	return nil
}
