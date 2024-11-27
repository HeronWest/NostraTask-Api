package config

import (
	"github.com/HeronWest/nostrataskapi/internal/auth"
	"github.com/HeronWest/nostrataskapi/internal/task"
	"github.com/HeronWest/nostrataskapi/internal/user"
)

type ApplicationBindings struct {
	di *DependencyInjector
}

func NewApplicationBindings(di *DependencyInjector) *ApplicationBindings {
	return &ApplicationBindings{
		di: di,
	}
}

func (ab *ApplicationBindings) InitializeBindings() error {
	// Initializing User
	ab.di.Provide(user.NewUserRepository)
	ab.di.Provide(user.NewUserService)
	ab.di.Provide(user.NewUserController)

	// Initializing Auth
	ab.di.Provide(auth.NewAuthRepository)
	ab.di.Provide(auth.NewAuthService)
	ab.di.Provide(auth.NewAuthController)

	ab.di.Provide(task.NewTaskRepository)
	ab.di.Provide(task.NewTaskService)
	ab.di.Provide(task.NewTaskController)

	return nil
}
