package classfile

/*
	JVM规范的字段结构定义

		field_info {
			u2				access_flags;
			u2				name_index;
			u2				descriptor_index;
			u2				attributes_count;
			attribute_info 	attributes[attributes_count];
		}
*/

// 标识字段信息和方法信息
type MemberInfo struct {
	cp ConstantPool
	// 字段或方法访问标志，u2类型
	accessFlags uint16
	// 常量池索引，指向字段或方法名，u2类型
	nameIndex uint16
	// 常量池索引，指向字段或方法的描述符，u2类型
	descriptorIndex uint16
	// 属性表
	attributes []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或方法信息
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// getter
func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

// 从常量池查找字段或方法名
func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

// 从常量池查找字段或方法描述符
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}
