package rtda

import "math"

type LocalVars [] Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) setInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) getInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) setFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) getFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self LocalVars) setLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index + 1]. num = int32(val >> 32)
}

func (self LocalVars) getLong(index uint) int64 {
	low := uint32(self[index]. num)
	high := uint32(self[index + 1]. num)
	return int64(high) << 32 | int64(low)
}