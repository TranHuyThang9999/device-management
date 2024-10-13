package fxloader

import (
	"device_management/api/controller"
	"device_management/api/middleware"
	"device_management/api/routers"
	"device_management/core/pgsql"
	"device_management/core/pgsql/repos"
	"device_management/core/usecase"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadValidator()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecase.NewUseCaseUser),
		fx.Provide(usecase.NewUseCaseJwt),
		fx.Provide(usecase.NewUseCaseDevice),
		fx.Provide(usecase.NewUseCaseFileStore),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(controller.NewControllerSaveFile),
		fx.Provide(controller.NewControllerUser),
		fx.Provide(controller.NewBaseController),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(controller.NewControllerDevices),
		fx.Provide(controller.NewControllerFileStore),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(pgsql.ConnectPgsql),
		fx.Provide(repos.NewUserRepository),
		fx.Provide(repos.NewDevicesRepository),
		fx.Provide(repos.NewFileStoreRepository),
		fx.Provide(repos.NewDBHelper),
	}
}
