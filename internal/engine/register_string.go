// Code generated by "stringer -type=Register -linecomment"; DO NOT EDIT.

package engine

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[A-0]
	_ = x[B-1]
	_ = x[C-2]
	_ = x[D-3]
	_ = x[E-4]
	_ = x[F-5]
	_ = x[H-6]
	_ = x[L-7]
	_ = x[AF-8]
	_ = x[BC-9]
	_ = x[DE-10]
	_ = x[HL-11]
}

const _Register_name = "ABCDEFHLAFBCDEHL"

var _Register_index = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 16}

func (i Register) String() string {
	if i < 0 || i >= Register(len(_Register_index)-1) {
		return "Register(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Register_name[_Register_index[i]:_Register_index[i+1]]
}