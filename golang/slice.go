package golang

type Slice struct {
	Type
	ValueType IType
}

func MakeSliceType(isPointer bool, valueType IType) Slice {
	return Slice{
		Type: Type{
			IsPointer: isPointer,
			IsSlice:   true,
		},
		ValueType: valueType,
	}
}
