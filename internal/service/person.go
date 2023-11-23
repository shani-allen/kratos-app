package service

import (
	"context"
	"math/rand"
	"strconv"

	v1 "basic-kratos-app/api/person/v1"
	"basic-kratos-app/internal/biz"
)

// PersonService is a greeter service.
type PersonService struct {
	v1.UnimplementedPersonServiceServer

	uc *biz.PersonUsecase
}

type PersonServicer interface{
	CreateGreeter(ctx context.Context, in *v1.GetPersonIdRequest) (*v1.Person, error)
}
// NewPersonService new a greeter service.
func NewGreeterService(uc *biz.PersonUsecase) *PersonService {
	return &PersonService{uc: uc}
}

// get person by id.
func (s *PersonService) GetPersonById(ctx context.Context, in *v1.GetPersonIdRequest) (*v1.Person, error) {
	g, err := s.uc.GetPersonById(ctx, in)
	if err != nil {
		return nil, err
	}
	return g , nil
}

// create person 
func(s *PersonService)CreatePerson(ctx context.Context, p *v1.CreatePersonRequest)(*v1.Person,error){
	id:=rand.Int()
	person, err:=s.uc.CreatePerson(ctx,&v1.Person{Name:p.Name,Email: p.Email, PersonId: strconv.Itoa(id)})
	if err!=nil{
		return nil, err
	}

	return &v1.Person{Name: person.Name, Email: person.Email,PersonId: strconv.Itoa(id)}, nil
}

// update the person detail
func(s *PersonService)UpdatePerson(ctx context.Context, p *v1.UpdatePersonRequest)(*v1.Person,error){
	//to do call method of handler layer
	return s.uc.UpdatePerson(ctx,&v1.Person{Name:p.Name,Email: p.Email, PersonId: p.PersonId})
}

// delete person by id 
func(s *PersonService)DeletePerson(ctx context.Context, id *v1.DeletePersonRequest)(*v1.DeletePersonResponse, error){
	return s.uc.DeletePerson(ctx,id)
}

