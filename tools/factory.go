package tools

import (
	"errors"
	"github.com/aerostatka/third-party-integrations/config"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/aerostatka/third-party-integrations/oktatool"
)

type Factory interface {
	Create(params *models.FactoryParams) (*Tool, error)
}

type Tool interface {
	PerformAction() *models.ActionResult
}

type ConsoleToolsFactory struct {
	appConfig *config.AppConfig
}

func CreateConsoleToolsFactory(conf *config.AppConfig) *ConsoleToolsFactory {
	return &ConsoleToolsFactory{
		appConfig: conf,
	}
}

func (factory *ConsoleToolsFactory) Create(params *models.FactoryParams) (Tool, error) {
	switch params.ToolType {
	case models.ToolTypeOkta:
		return oktatool.CreateOktaTool(factory.appConfig, params.ActionType, params.ActionParameters)
	}

	return nil, errors.New("Action cannot be created")
}
