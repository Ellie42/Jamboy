// Code generated by "stringer -type ValueKeyword"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValKeywordZ-0]
	_ = x[ValKeywordNZ-1]
	_ = x[ValKeywordCB-2]
}

const _ValueKeyword_name = "ValKeywordZValKeywordNZValKeywordCB"

var _ValueKeyword_index = [...]uint8{0, 11, 23, 35}

func (i ValueKeyword) String() string {
	if i < 0 || i >= ValueKeyword(len(_ValueKeyword_index)-1) {
		return "ValueKeyword(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueKeyword_name[_ValueKeyword_index[i]:_ValueKeyword_index[i+1]]
}
