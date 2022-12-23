package entity

type EntityValidator interface {
	Validate() error
}

func Validate(entity EntityValidator) error {
	return entity.Validate()
}
