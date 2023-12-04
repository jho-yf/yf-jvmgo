package classfile

type ConstantLongInfo struct {
	val int64
}

func (longInfo *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	longInfo.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (doubleInfo *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	doubleInfo.val = float64(bytes)
}
