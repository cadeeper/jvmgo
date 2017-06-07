package classfile

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
