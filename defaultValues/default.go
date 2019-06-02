package defaultValues

type ReturnType string
type FunctionConfig func(ReturnType) (ReturnType, error)

func A(a ReturnType) (ReturnType, error) {
	return a + " done", nil
}

type Thing struct {
	functionConfig FunctionConfig
	returnType     ReturnType
}

func NewThing(options ...func(*Thing) error) (ReturnType, error) {
	f := &Thing{}

	f.functionConfig = func(b ReturnType) (ReturnType, error) { return b + "  default", nil }
	f.returnType = " ...some default..."

	for _, op := range options {
		err := op(f)
		if err != nil {
			return "", err
		}
	}
	return f.functionConfig(f.returnType)
}

func OptionalFn(f *Thing) error {
	f.functionConfig = A
	return nil
}

func OptionalReturnType(t ReturnType) func(f *Thing) error {
	return func(f *Thing) error {
		f.returnType = t
		return nil
	}
}
