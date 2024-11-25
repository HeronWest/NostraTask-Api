package config

import "go.uber.org/dig"

type DependencyInjector struct {
	container *dig.Container
	logger    *Logger
}

func NewDependencyInjector() *DependencyInjector {
	return &DependencyInjector{
		container: dig.New(),
		logger:    GetLogger("$dependency_injector: "),
	}
}

func (di *DependencyInjector) Provide(constructor interface{}) {
	err := di.container.Provide(constructor)
	if err != nil {
		di.logger.Errorf("Erro ao prover dependência: %v", err)
		return
	}
}

func (di *DependencyInjector) Invoke(function interface{}) {
	err := di.container.Invoke(function)
	if err != nil {
		di.logger.Errorf("Erro ao invocar função: %v", err)
		return
	}
}
