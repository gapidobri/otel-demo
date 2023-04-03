package service

type (
	Service struct{}
)

func NewService() Service {
	return Service{}
}

func (s Service) Get() string {
	return "Hello World!"
}
