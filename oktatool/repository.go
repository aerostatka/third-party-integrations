package oktatool

import (
	"context"
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
	GetApplications(status string, hardLimit int) ([]models.SimpleApp, error)
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
