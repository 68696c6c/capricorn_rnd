package golang

import "fmt"

type Slice struct {
	*Type
	isPointer bool
}

func (s *Slice) GetType() *Type {
	return s.Type
}

func (s *Slice) CopyType() *Type {
	return copyType(s.Type)
}

func (s *Slice) GetReference() string {
	prefix := ""
	if s.isPointer {
		prefix = "*"
	}
	return fmt.Sprintf("%s%s", prefix, s.Type.GetReference())
}

func MakeSliceType(isPointer bool, valueType IType) *Slice {
	v := valueType.CopyType()
	v.IsSlice = true
	return &Slice{
		Type:      v,
		isPointer: isPointer,
	}
}
