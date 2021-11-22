package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type ListAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateListAppAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *ListAppAction {
	return &ListAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *ListAppAction) ApplyAction() *models.ActionResult {
	_, err := action.listApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Application list is stored in %s", action.parameters.DataLocation),
	)
}

func (action *ListAppAction) listApplications() ([]models.SimpleApp, error) {
	action.logger.Info("Fetching applications....")
	apps, err := action.repository.GetApplications(action.parameters.Status, action.parameters.Limit)

	if err != nil {
		return apps, err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications were found", len(apps)))

	action.logger.Info("Storing applications to a file")
	err = action.storage.StoreApplicationsData(action.parameters.DataLocation, apps)

	if err != nil {
		return apps, err
	}

	action.logger.Info("Done.")

	return apps, nil
}
