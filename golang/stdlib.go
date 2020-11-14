package golang

func MakeErrorType() Type {
	return MockType("", "error", false, false)
}

func MakeTimeType(isPointer bool) Type {
	return MockType("time", "Time", isPointer, false)
}
