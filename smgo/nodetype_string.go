// Code generated by "stringer -type=NodeType"; DO NOT EDIT.

package smgo

import "strconv"

const _NodeType_name = "PackageNodeFunctionNodeFieldNodeImportNodeConstNodeVarNodeTypeNodeStructNodeInterfaceNodeComment"

var _NodeType_index = [...]uint8{0, 11, 23, 32, 42, 51, 58, 66, 76, 89, 96}

func (i NodeType) String() string {
	if i < 0 || i >= NodeType(len(_NodeType_index)-1) {
		return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeType_name[_NodeType_index[i]:_NodeType_index[i+1]]
}
