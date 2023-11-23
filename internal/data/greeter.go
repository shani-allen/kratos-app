package data

import (
	"context"

	v1 "basic-kratos-app/api/helloworld/v1"
	"basic-kratos-app/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type personRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.PersonRepo {
	return &personRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *personRepo) Save(ctx context.Context, p *v1.Person) (*biz.Person, error) {
	return &biz.Person{Name:p.Name, Email:p.Email}, nil
}

func (r *personRepo) Update(ctx context.Context, p *v1.Person) (*biz.Person, error) {
	return  &biz.Person{Name:p.Name, Email:p.Email}, nil
}

func (r *personRepo) FindByID(ctx context.Context, id string) (*biz.Person, error) {
	return &biz.Person{Name:"shani", Email:"shani.123@gamil.com"}, nil
}

func (r *personRepo) ListByHello(ctx context.Context, id []string) ([]*biz.Person, error) {
	return nil, nil
}

func (r *personRepo) ListAll(context.Context) ([]*biz.Person, error) {
	return nil, nil
}
