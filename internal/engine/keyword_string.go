// Code generated by "stringer -type Keyword -linecomment"; DO NOT EDIT.

package engine

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KeywordZ-0]
	_ = x[KeywordNZ-1]
	_ = x[KeywordCB-2]
	_ = x[KeywordC-3]
	_ = x[KeywordNC-4]
}

const _Keyword_name = "ZNZCBCNC"

var _Keyword_index = [...]uint8{0, 1, 3, 5, 6, 8}

func (i Keyword) String() string {
	if i < 0 || i >= Keyword(len(_Keyword_index)-1) {
		return "Keyword(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Keyword_name[_Keyword_index[i]:_Keyword_index[i+1]]
}
