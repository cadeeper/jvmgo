package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
