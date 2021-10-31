package application

import (
	"github.com/aerostatka/third-party-integrations/config"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/tools"
)

func Start(params []string)  {
	log := logger.CreateNewZapLogger()

	log.Info("Application start")

	appConfig, err := config.CreateContextConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Config has been parsed successfully")

	actionFactory := tools.CreateConsoleToolsFactory(appConfig.GetConfig())
	factoryParams, err := tools.ConvertToFactoryParams(params)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = factoryParams.Validate()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Config has been validated successfully")

	_, err = actionFactory.Create(factoryParams)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Action has been created successfully")


	log.Info("Application finish")
}