// Code generated by "stringer -type=ContainerType"; DO NOT EDIT.

package smgo

import "strconv"

const _ContainerType_name = "StructContainer"

var _ContainerType_index = [...]uint8{0, 15}

func (i ContainerType) String() string {
	if i < 0 || i >= ContainerType(len(_ContainerType_index)-1) {
		return "ContainerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ContainerType_name[_ContainerType_index[i]:_ContainerType_index[i+1]]
}
