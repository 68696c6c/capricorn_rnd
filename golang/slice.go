package golang

import "fmt"

type Slice struct {
	*Type
	valueType IType
}

func (s Slice) GetReference() string {
	prefix := "[]"
	if s.IsPointer {
		prefix = "*[]"
	}
	return fmt.Sprintf("%s%s", prefix, s.valueType.GetReference())
}

func MakeSliceType(isPointer bool, valueType IType) *Slice {
	return &Slice{
		Type:      MockType(valueType.GetImport(), valueType.GetName(), isPointer, true),
		valueType: valueType,
	}
}
