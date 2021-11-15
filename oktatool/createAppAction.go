package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type CreateAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateCreateAppAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *CreateAppAction {
	return &CreateAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *CreateAppAction) ApplyAction() *models.ActionResult {
	err := action.createApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were created", action.parameters.DataLocation),
	)
}

func (action *CreateAppAction) createApplications() error {
	action.logger.Info("Fetching applications from storage....")
	apps, err := action.storage.GetAppsData(action.parameters.DataLocation)
	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications successfully fetched.", len(apps)))
	action.logger.Info("Fetching creation templates...")

	template, err := action.storage.GetTemplate(action.parameters.TemplateLocation, action.parameters.TemplateName)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%s template has been found.", template.Name))
	action.logger.Info("Fetching applications from OKTA....")

	activeApps, err := action.repository.GetApplications(
		models.OktaParamsApplicationsStatusActive,
		action.parameters.Limit,
	)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%d active applications were found in OKTA.", len(activeApps)))

	for _, app := range apps {
		action.logger.Info(fmt.Sprintf("Processing %s application.", app.Label))
		oktaApp := app.FindStorageAppInList(activeApps)

		if oktaApp == nil {
			action.logger.Info(fmt.Sprintf("Application %s is not found in OKTA.", app.Label))

			err = app.ValidateForCreation()
			if err != nil {
				action.logger.Info(
					fmt.Sprintf("Error occurred during processing application: %s.", err.Error()),
				)
				action.logger.Info("Processing skipped.")
			}
		} else {
			action.logger.Info("Active application %s is found in OKTA. Application cannot be added.")
		}
	}

	return nil
}
