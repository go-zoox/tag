package attribute

import "strings"

// isScalarPointer reports whether typ is a pointer to a scalar (not slice, struct, or map).
func isScalarPointer(typ string) bool {
	if !strings.HasPrefix(typ, "*") {
		return false
	}
	elem := typ[1:]
	if strings.HasPrefix(elem, "[]") || strings.Contains(elem, "struct") {
		return false
	}
	return true
}

func baseType(typ string) string {
	if isScalarPointer(typ) {
		return typ[1:]
	}
	return typ
}

func matchesScalarType(typ, scalar string) bool {
	return typ == scalar || typ == "*"+scalar
}

func isIntBaseType(bt string) bool {
	switch bt {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	default:
		return false
	}
}

func isFloatBaseType(bt string) bool {
	switch bt {
	case "float", "float32", "float64":
		return true
	default:
		return false
	}
}

func isBoolType(typ string) bool {
	return matchesScalarType(typ, "bool")
}
