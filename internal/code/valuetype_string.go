// Code generated by "stringer -type ValueType"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValTypeRegister-0]
	_ = x[ValTypeRead-1]
	_ = x[ValTypeConst-2]
	_ = x[ValTypeKeyword-3]
}

const _ValueType_name = "ValTypeRegisterValTypeReadValTypeConstValTypeKeyword"

var _ValueType_index = [...]uint8{0, 15, 26, 38, 52}

func (i ValueType) String() string {
	if i < 0 || i >= ValueType(len(_ValueType_index)-1) {
		return "ValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueType_name[_ValueType_index[i]:_ValueType_index[i+1]]
}
