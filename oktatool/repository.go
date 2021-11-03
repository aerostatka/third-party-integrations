package oktatool

import (
	"context"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

const (
	OktaRepositoryApplicationsListLimit = 50
	OktaRepositoryApplicationsHardLimit = 1000
)

type Repository interface {
	GetApplications() ([]models.SimpleApp, error)
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

func (rep *OktaRepository) GetApplications() ([]models.SimpleApp, error) {
	var applications []okta.App
	var apps []models.SimpleApp
	var resp *okta.Response
	var err error

	for {
		if resp == nil {
			applications, resp, err = rep.client.Application.ListApplications(rep.context, &query.Params{
				Limit: OktaRepositoryApplicationsListLimit,
			})
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

		if len(apps) >= OktaRepositoryApplicationsHardLimit {
			break
		}

		if !resp.HasNextPage() {
			break
		}
	}

	return apps, nil
}
