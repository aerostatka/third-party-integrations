package tools

import "github.com/aerostatka/third-party-integrations/config"

type Factory interface {
	Create(params *FactoryParams) (*Tool, error)
}

type Tool interface {
	PerformAction() (ActionResult, error)
}

type ConsoleToolsFactory struct {
	appConfig *config.AppConfig
}

func (factory *ConsoleToolsFactory) Create(params *FactoryParams) (*Tool, error)  {
	return nil, nil
}

func CreateConsoleToolsFactory(conf *config.AppConfig) *ConsoleToolsFactory  {
	return &ConsoleToolsFactory{
		appConfig: conf,
	}
}
