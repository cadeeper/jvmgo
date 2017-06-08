package classfile


/**
	CONSTANT_Fieldref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}

	CONSTANT_Methodref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}

	CONSTANT_InterfaceMethodref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
	Fields, methods, and interface methods are represented by similar structure
 */
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16			//类在常量池中的索引
	nameAndTypeIndex uint16			//名字和类型在常量池中的索引
}

/**
	字段
 */
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

/**
	方法
 */
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

/**
	接口方法
 */
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
