package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	OktaParamsDefaultTemplateLocation    = "/data/templates/applications/default.json"
	OktaParamsDefaultTemplateName        = "default"
	OktaParamsApplicationsStatusActive   = "ACTIVE"
	OktaParamsApplicationsStatusInactive = "INACTIVE"
)

var (
	OktaParamsStatuses = map[string]bool{
		OktaParamsApplicationsStatusActive:   true,
		OktaParamsApplicationsStatusInactive: true,
	}
)

type OktaToolParameters struct {
	Action           string
	Name             string
	DataLocation     string
	TemplateLocation string
	TemplateName     string
	CertLocation     string
	Limit            int
	Status           string
}

func (toolParams *OktaToolParameters) LoadParameters(params []string) error {
	switch toolParams.Action {
	case ActionTypeOktaApplicationCreate:
		toolParams.DataLocation = params[0]

		length := len(params)
		if length > 1 {
			toolParams.TemplateName = params[1]
		} else {
			toolParams.TemplateName = OktaParamsDefaultTemplateName
		}

		if length > 2 {
			toolParams.TemplateLocation = params[2]
		} else {
			path, err := os.Getwd()

			if err != nil {
				return errors.New("Error finding current directory")
			}

			fullPath := path + OktaParamsDefaultTemplateLocation
			toolParams.TemplateLocation = fullPath
		}

		if length > 3 {
			toolParams.CertLocation = params[3]
		}
	case ActionTypeOktaApplicationList:
		toolParams.DataLocation = params[0]
		toolParams.Status = ""

		length := len(params)
		if length > 1 {
			limit, err := strconv.Atoi(params[1])

			if err != nil {
				return err
			}

			toolParams.Limit = limit
		}

		if length > 2 {
			if !OktaParamsStatuses[params[2]] {
				return errors.New("Provided status is not valid")
			}

			toolParams.Status = params[2]
		}
	case ActionTypeOktaApplicationGet:
		toolParams.DataLocation = params[0]
		toolParams.Name = params[1]
	case ActionTypeOktaApplicationDisable,
		ActionTypeOktaApplicationEnable,
		ActionTypeOktaApplicationDelete:
		toolParams.DataLocation = params[0]
	case ActionTypeOktaApplicationUpdateCert:
		toolParams.DataLocation = params[0]
		toolParams.CertLocation = params[1]
	default:
		return errors.New(fmt.Sprintf("Action %s is not supported", toolParams.Action))
	}

	return nil
}
