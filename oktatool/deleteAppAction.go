package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type DeleteAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateDeleteAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *DeleteAppAction {
	return &DeleteAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *DeleteAppAction) ApplyAction() *models.ActionResult {
	err := action.disableApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were deleted", action.parameters.DataLocation),
	)
}

func (action *DeleteAppAction) disableApplications() error {
	action.logger.Info("Fetching applications from storage....")
	apps, err := action.storage.GetAppsData(action.parameters.DataLocation)
	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications successfully fetched", len(apps)))
	action.logger.Info("Fetching applications from OKTA....")

	activeApps, err := action.repository.GetApplications(
		"",
		action.parameters.Limit,
	)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%d applications were found", len(activeApps)))

	for _, app := range apps {
		action.logger.Info(fmt.Sprintf("Processing %s application.", app.Label))
		oktaApp := app.FindStorageAppInList(activeApps)

		if oktaApp == nil {
			action.logger.Info(fmt.Sprintf("Application %s is not found in OKTA.", app.Label))
		} else {
			action.logger.Info(fmt.Sprintf(
				"Application %s is found in OKTA with the status %s.",
				oktaApp.Label,
				oktaApp.Status),
			)

			if oktaApp.Status == models.OktaParamsApplicationsStatusActive {
				action.logger.Info("Disabling application....")
				err := action.repository.ChangeApplicationStatus(oktaApp, models.OktaParamsApplicationsStatusInactive)
				if err != nil {
					return err
				}
				action.logger.Info("Done.")
			}

			action.logger.Info("Deleting application....")
			err := action.repository.DeleteApplication(oktaApp)
			if err != nil {
				return err
			}
			action.logger.Info("Done.")
		}
	}

	return nil
}
