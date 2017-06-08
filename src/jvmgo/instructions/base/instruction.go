package base

import "jvmgo/rtda"

type Instruction interface {

	FetchOperands(reader *BytecodeReader)

	Execute(frame *rtda.Frame)

}

/**
	没有操作数的指令
 */
type NoOperandsInstruction struct {

}

/**
	跳转指令
 */
type BranchInstruction struct {
	Offset int		//跳转偏移量
}

/**
	存储和加载类指令
 */
type Index8Instruction struct {
	Index uint		//局部变量表索引
}

type Index16Instruction struct {
	Index uint		//常量池索引
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}