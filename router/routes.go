package router

import (
	"github.com/HeronWest/nostrataskapi/config"
	"github.com/HeronWest/nostrataskapi/internal/user"
)

func initializeRoutes(di *config.DependencyInjector) {
	di.Invoke(user.SetupRoutes)
}
