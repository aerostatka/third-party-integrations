package models

import (
	"errors"
	"os"
	"strconv"
)

const (
	OktaParamsDefaultTemplateLocation = "/data/templates/application/default.json"
	OktaParamsActiveStatus            = "active"
)

type OktaToolParameters struct {
	Action           string
	DataLocation     string
	TemplateLocation string
	TemplateName     string
	CertLocation     string
	Limit            int
	OnlyActive       bool
}

func (toolParams *OktaToolParameters) LoadParameters(params []string) error {
	switch toolParams.Action {
	case ActionTypeOktaApplicationCreate:
		toolParams.DataLocation = params[0]
		toolParams.TemplateName = params[1]

		length := len(params)
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
			toolParams.TemplateLocation = params[3]
		}
	case ActionTypeOktaApplicationList:
		toolParams.DataLocation = params[0]
		toolParams.OnlyActive = false

		length := len(params)
		if length > 1 {
			limit, err := strconv.Atoi(params[1])

			if err != nil {
				return err
			}

			toolParams.Limit = limit
		}

		if length > 2 {
			if params[2] == OktaParamsActiveStatus {
				toolParams.OnlyActive = true
			} else {
				return errors.New("Third parameter is incorrect")
			}
		}
	case ActionTypeOktaApplicationDisable,
		ActionTypeOktaApplicationEnable,
		ActionTypeOktaApplicationDelete:
		toolParams.DataLocation = params[0]
	case ActionTypeOktaApplicationUpdateCert:
		toolParams.DataLocation = params[0]
		toolParams.CertLocation = params[1]
	default:
		return errors.New("Action " + toolParams.Action + " is not supported")
	}

	return nil
}
