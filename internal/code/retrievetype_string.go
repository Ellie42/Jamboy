// Code generated by "stringer -type RetrieveType"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RetrieveVal-0]
	_ = x[RetrievePointer-1]
}

const _RetrieveType_name = "RetrieveValRetrievePointer"

var _RetrieveType_index = [...]uint8{0, 11, 26}

func (i RetrieveType) String() string {
	if i < 0 || i >= RetrieveType(len(_RetrieveType_index)-1) {
		return "RetrieveType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RetrieveType_name[_RetrieveType_index[i]:_RetrieveType_index[i+1]]
}
