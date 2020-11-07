package generator

type ErrorHandler interface {
	HandleError(e error)
}

type PanicHandler struct{}

func (h PanicHandler) HandleError(e error) {
	if e != nil {
		panic(e)
	}
}
