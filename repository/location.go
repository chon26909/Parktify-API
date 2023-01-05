package repository

type Location interface{}

type location struct{}

func NewLocationRepository() Location {
	return &location{}
}

func (l *location) GetAllLocation() ([]Location, error) {
	return nil, nil
}
