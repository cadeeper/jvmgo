package base

type BytecodeReader struct {
	code [] byte
	pc   int
}

func (self *BytecodeReader) Reset(code [] byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	top := uint16(self.ReadUint8())
	low := uint16(self.ReadUint8())
	return (top << 8) | low
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadUint32() uint32 {
	byte1 := uint32(self.ReadUint8())
	byte2 := uint32(self.ReadUint8())
	byte3 := uint32(self.ReadUint8())
	byte4 := uint32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (self *BytecodeReader) ReadInt32() int32 {
	return int32(self.ReadUint32())
}