package validation

type Validation interface {
	Validation(request interface{}) error
}
