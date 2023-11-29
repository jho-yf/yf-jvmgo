package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// 读取u1类型数据
//
// u1类型: 1字节无符号整数
func (cr *ClassReader) readUint8() uint8 {
	value := cr.data[0]
	cr.data = cr.data[1:]
	return value
}

// 读取u2类型数据
//
// u2类型：2字节无符号整数
func (cr *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return value
}

// 读取u4类型数据
//
// u4类型：4字节无符号整数
func (cr *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return value
}

// 读取uint64类型数据
//
// JVM规范没有定义u8
func (cr *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return value
}

// 读取uint16表，表的大小由开头的uint16数据定义
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s
}

// 
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
