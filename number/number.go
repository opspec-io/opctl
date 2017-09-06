// Package number implements usecases surrounding numbers
package number

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ Number

type Number interface {
	interpreter
	validator
}

func New() Number {
	return _Number{
		interpreter: newInterpreter(),
		validator:   newValidator(),
	}
}

type _Number struct {
	interpreter
	validator
}
