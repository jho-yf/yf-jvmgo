package classfile

import "fmt"

// 定义JVM规范的class文件格式
type ClassFile struct {
	// 魔数（magic number）: u4类型，即4字节无符号整数
	magic uint32
	// 次版本：u2类型，即2字节无符号整数
	minorVersion uint16
	// 主版本：u2类型，即2字节无符号整数
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标记：16位的bitmask，指出class文件定义的是类还是接口，访问级别是private还是public
	accessFlags uint16
	// 常量池索引：u2类型，指向当前类名
	thisClass uint16
	// 常量池索引：u2类型，指向当前父类名
	superClass uint16
	// 常量池索引：u2类型，指向当前接口
	interfaces []uint16
	// 字段表
	fields []*MemberInfo
	// 方法表
	methods    []*MemberInfo
	attributes []AttributeInfo
}

// 将[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		// panic-recover机制。see https://zhuanlan.zhihu.com/p/660174941
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read方法一次调用其他方法解析class文件
func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

// 魔数 magic number
// 规定满足改格式的文件必须以固定几个字节开头，这几个字节主要起标识作用
// 例如：
// PDF：以"%PDF"，即0x25、0x50、0x44、0x46，4字节开头
// ZIP：以"PK"，即0x25、0x4B，2字节开头
// class文件：以"0xCAFEBABE"开头
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFBABE {
		// JVM规范，如果加载class文件不符合要求的格式，抛出"java.lang.ClassFormatError"异常
		panic("java.lang.ClassFormatError: magic!")
	}
	cf.magic = magic
}

// 特定JVM实现只能支持版本号在某个范围内的class文件
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	// 如果遇到不支持的class文件版本号则，抛出"java.lang.UnsupportedClassVersionError"异常
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// getter
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// getter
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// getter
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

// getter
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

// getter
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

// 从常量池中查找父类名
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return cf.constantPool.getClassName(cf.superClass)
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
