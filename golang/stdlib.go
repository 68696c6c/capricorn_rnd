package golang

func MakeTypeError() *Type {
	return MockType("", "error", false, false)
}

func MakeTypeTime(isPointer bool) *Type {
	return MockType("time", "Time", isPointer, false)
}

func MakeTypeInt(isPointer bool) *Type {
	return MockType("", "int", isPointer, false)
}

func MakeTypeInterfaceLiteral() *Type {
	return MockType("", "interface{}", false, false)
}

func MakeTypeDriverValue() *Type {
	return MockType("database/sql/driver", "Value", false, false)
}

func MakeTypeString(isPointer bool) *Type {
	return MockType("", "string", isPointer, false)
}
