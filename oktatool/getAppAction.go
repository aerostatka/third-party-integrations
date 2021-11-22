package oktatool

import (
	"errors"
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type GetAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateGetAppAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *GetAppAction {
	return &GetAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *GetAppAction) ApplyAction() *models.ActionResult {
	err := action.getApplicationData()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were created", action.parameters.DataLocation),
	)
}

func (action *GetAppAction) getApplicationData() error {
	action.logger.Info("Fetching applications....")
	apps, err := action.repository.GetApplications(models.OktaParamsApplicationsStatusActive, action.parameters.Limit)

	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d active applications were found", len(apps)))

	var foundApp *models.SimpleApp
	for _, app := range apps {
		if app.Label == action.parameters.Name {
			foundApp = &app
			break
		}
	}

	if foundApp != nil {
		application, err := action.repository.GetApplication(foundApp.Id)
		if err != nil {
			return err
		}

		action.logger.Info(
			fmt.Sprintf("Application %s data has been retrieved from OKTA", foundApp.Label),
		)

		action.storage.StoreApplicationData(action.parameters.DataLocation, application)
		if err != nil {
			return err
		}

		action.logger.Info(
			fmt.Sprintf("Data has been stored to file %s", action.parameters.DataLocation),
		)
	} else {
		return errors.New("Application is not found")
	}

	return nil
}
