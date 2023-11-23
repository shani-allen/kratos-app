package data

import (
	"context"
	"errors"


	v1 "basic-kratos-app/api/person/v1"
	"basic-kratos-app/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type personRepo struct {
	data *Data
	log  *log.Helper
	persons map[string]v1.Person
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.PersonRepo {
	return &personRepo{
		data: data,
		log:  log.NewHelper(logger),
		persons: make(map[string]v1.Person),
	}
}

//for now i am using in memory just for veryfication 

func (r *personRepo) Save(ctx context.Context, p *v1.Person) (*v1.Person, error) {
	id:=p.PersonId
	_, ok:=r.persons[id]
	if ok{
		return nil,errors.New("already person exists in DB") 
	}
	r.persons[id]=*p
	return p, nil
}

func (r *personRepo) Update(ctx context.Context, p *v1.Person) (*v1.Person, error) {
	 id:=p.PersonId
	_,ok:=r.persons[id]
	if ok{
		r.persons[id]=*p
	}
	
	return nil, errors.New("person does not exists")
}

func (r *personRepo) FindByID(ctx context.Context, id string) (*v1.Person, error) {
	p,ok:=r.persons[id]
	if ok{
		return &p,nil
	}

	return nil, errors.New("person does not exists")
}

func(r *personRepo)DeleteById(ctx context.Context, id *v1.DeletePersonRequest)(*v1.DeletePersonResponse, error){
	_,ok:=r.persons[id.PersonId]
	if ok{
		delete(r.persons,id.PersonId)
		return &v1.DeletePersonResponse{Success: true}, nil
	}

	return nil, errors.New("person does not exists")
}
