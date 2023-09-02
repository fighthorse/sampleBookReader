//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/fighthorse/sampleBookReader/gin-admin/internal/app/api"
	"github.com/fighthorse/sampleBookReader/gin-admin/internal/app/dao"
	"github.com/fighthorse/sampleBookReader/gin-admin/internal/app/module/adapter"
	"github.com/fighthorse/sampleBookReader/gin-admin/internal/app/router"
	"github.com/fighthorse/sampleBookReader/gin-admin/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
