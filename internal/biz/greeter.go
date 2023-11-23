package biz

import (
	"context"

	v1 "basic-kratos-app/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Person struct {
	Name string
	Email string

}

// GreeterRepo is a Greater repo.
type PersonRepo interface {
	Save(context.Context, *v1.Person) (*Person, error)
	Update(context.Context, *v1.Person) (*Person, error)
	FindByID(context.Context, string) (*Person, error)
	ListByHello(context.Context, []string) ([]*Person, error)
	ListAll(context.Context) ([]*Person, error)
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
func (uc *PersonUsecase) CreateGreeter(ctx context.Context, in *v1.GetPersonIdRequest) (*Person, error) {
	uc.log.WithContext(ctx).Infof("Name: %v, Email: %v", in.PersonId)
	return uc.repo.FindByID(ctx, in.PersonId)
}
