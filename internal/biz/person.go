package biz

import (
	"context"

	v1 "basic-kratos-app/api/person/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)


// GreeterRepo is a Greater repo.
type PersonRepo interface {
	Save(context.Context, *v1.Person) (*v1.Person, error)
	Update(context.Context, *v1.Person) (*v1.Person, error)
	FindByID(context.Context, string) (*v1.Person, error)
	DeleteById(context.Context, *v1.DeletePersonRequest)(*v1.DeletePersonResponse ,error)
}

// GreeterUsecase is a Greeter usecase.
type PersonUsecase struct {
	repo PersonRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo PersonRepo, logger log.Logger) *PersonUsecase {
	return &PersonUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *PersonUsecase) GetPersonById(ctx context.Context, in *v1.GetPersonIdRequest) (*v1.Person, error) {
	uc.log.WithContext(ctx).Infof("Name: %v, Email: %v", in.PersonId)
	return uc.repo.FindByID(ctx, in.PersonId)
}

// CreatePerson in the DB
func(uc *PersonUsecase)CreatePerson(ctx context.Context, p *v1.Person)(*v1.Person,error){
	uc.log.WithContext(ctx).Infof("person: %v",p)
	return uc.repo.Save(ctx,p)
}

//
func(uc *PersonUsecase)UpdatePerson(ctx context.Context, p *v1.Person)(*v1.Person,error){
	uc.log.WithContext(ctx).Infof("person: %v",p)
	return  uc.repo.Update(ctx,p)
}

//
func(uc *PersonUsecase)DeletePerson(ctx context.Context, p *v1.DeletePersonRequest)(*v1.DeletePersonResponse,error){
	uc.log.WithContext(ctx).Infof("person: %v", p)
	return  uc.repo.DeleteById(ctx,p)
}