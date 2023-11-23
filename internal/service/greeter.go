package service

import (
	"context"

	v1 "basic-kratos-app/api/helloworld/v1"
	"basic-kratos-app/internal/biz"
)

// GreeterService is a greeter service.
type PersonService struct {
	v1.UnimplementedPersonServiceServer

	uc *biz.PersonUsecase
}

type PersonServicer interface{
	CreateGreeter(ctx context.Context, in *v1.GetPersonIdRequest) (*v1.Person, error)
}
// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.PersonUsecase) *PersonService {
	return &PersonService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *PersonService) GetPersonById(ctx context.Context, in *v1.GetPersonIdRequest) (*v1.Person, error) {
	g, err := s.uc.CreateGreeter(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.Person{Name: g.Name,Email: g.Email}, nil
}
