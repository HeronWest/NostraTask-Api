package config

import (
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

	return nil
}
