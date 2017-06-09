package rtda

type Frame struct {
	lower        *Frame
	thread       *Thread
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
	nextPC       int           //下一步操作
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}