package tools

import (
	"context"
	"errors"
	"github.com/aerostatka/third-party-integrations/config"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/aerostatka/third-party-integrations/oktatool"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type Factory interface {
	Create(params *models.FactoryParams) (*Tool, error)
}

type Tool interface {
	PerformAction() (*models.ActionResult, error)
}

type ConsoleToolsFactory struct {
	appConfig *config.AppConfig
}

func CreateConsoleToolsFactory(conf *config.AppConfig) *ConsoleToolsFactory {
	return &ConsoleToolsFactory{
		appConfig: conf,
	}
}

func createOktaTool(appConfig *config.AppConfig, action string, actionParameters []string) (*oktatool.ConsoleTool, error) {
	config := &models.OktaToolParameters{
		Action: action,
	}

	err := config.LoadParameters(actionParameters)

	if err != nil {
		return nil, err
	}

	ctx, client, err := okta.NewClient(
		context.Background(),
		okta.WithOrgUrl(appConfig.Okta.Domain),
		okta.WithToken(appConfig.Okta.Token),
	)

	if err != nil {
		return nil, err
	}

	return &oktatool.ConsoleTool{
		Context:    ctx,
		Client:     client,
		Parameters: config,
	}, nil
}

func (factory *ConsoleToolsFactory) Create(params *models.FactoryParams) (Tool, error) {
	switch params.ToolType {
	case models.ToolTypeOkta:
		return createOktaTool(factory.appConfig, params.ActionType, params.ActionParameters)
	}

	return nil, errors.New("Action cannot be created")
}
