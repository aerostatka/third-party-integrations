package oktatool

import (
	"context"
	"errors"
	"fmt"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

const (
	OktaRepositoryApplicationsListLimit = 50
	OktaRepositoryApplicationsHardLimit = 10000
)

type Repository interface {
	GetApplication(appId string) (okta.App, error)
	GetApplications(status string, hardLimit int) ([]models.SimpleApp, error)
	ChangeApplicationStatus(app *models.SimpleApp, status string) error
	DeleteApplication(app *models.SimpleApp) error
	CreateApplication(app *models.SimpleApp, template *models.Template) error
}

type OktaRepository struct {
	context context.Context
	client  *okta.Client
}

func CreateOktaRepository(ctx context.Context, ct *okta.Client) *OktaRepository {
	return &OktaRepository{
		context: ctx,
		client:  ct,
	}
}

func (rep *OktaRepository) GetApplication(appId string) (okta.App, error) {
	app, _, err := rep.client.Application.GetApplication(rep.context, appId, &models.OktaApplication{}, nil)

	return app, err
}

func (rep *OktaRepository) GetApplications(status string, hardLimit int) ([]models.SimpleApp, error) {
	var applications []okta.App
	var apps []models.SimpleApp
	var resp *okta.Response
	var err error

	if hardLimit <= 0 {
		hardLimit = OktaRepositoryApplicationsHardLimit
	}

	listLimit := OktaRepositoryApplicationsListLimit
	if hardLimit < listLimit {
		listLimit = hardLimit
	}

	for {
		if resp == nil {
			qr := &query.Params{
				Limit: int64(listLimit),
			}

			if status != "" {
				qr.Filter = fmt.Sprintf("status eq \"%s\"", status)
			}

			applications, resp, err = rep.client.Application.ListApplications(rep.context, qr)
		} else {
			resp, err = resp.Next(rep.context, &applications)
		}

		if err != nil {
			return apps, err
		}

		for _, a := range applications {
			app := models.SimpleApp{
				Id:     a.(*okta.Application).Id,
				Code:   a.(*okta.Application).Name,
				Label:  a.(*okta.Application).Label,
				Status: a.(*okta.Application).Status,
			}

			apps = append(apps, app)
		}

		if len(apps) >= hardLimit {
			break
		}

		if !resp.HasNextPage() {
			break
		}
	}

	return apps, nil
}

func (rep *OktaRepository) ChangeApplicationStatus(app *models.SimpleApp, status string) error {
	if app.Id == "" {
		return errors.New("Application id is not set up")
	}

	var err error
	if status == models.OktaParamsApplicationsStatusActive {
		_, err = rep.client.Application.ActivateApplication(rep.context, app.Id)
	} else {
		_, err = rep.client.Application.DeactivateApplication(rep.context, app.Id)
	}

	return err
}

func (rep *OktaRepository) DeleteApplication(app *models.SimpleApp) error {
	if app.Id == "" {
		return errors.New("Application id is not set up")
	}

	_, err := rep.client.Application.DeleteApplication(rep.context, app.Id)

	return err
}

func (rep *OktaRepository) CreateApplication(app *models.SimpleApp, template *models.Template) error {
	return nil
}
