// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"basic-kratos-app/internal/biz"
	"basic-kratos-app/internal/conf"
	"basic-kratos-app/internal/data"
	"basic-kratos-app/internal/server"
	"basic-kratos-app/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	personRepo := data.NewGreeterRepo(dataData, logger)
	personUsecase := biz.NewGreeterUsecase(personRepo, logger)
	personService := service.NewGreeterService(personUsecase)
	grpcServer := server.NewGRPCServer(confServer, personService, logger)
	httpServer := server.NewHTTPServer(confServer, personService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
