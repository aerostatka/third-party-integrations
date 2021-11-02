package models

import (
	"errors"
	"os"
)

const OktaParamsDefaultTemplateLocation = "/data/templates/application/default.json"

type OktaToolParameters struct {
	Action           string
	DataLocation     *string
	TemplateLocation *string
	TemplateName     *string
	CertLocation     *string
}

func (toolParams *OktaToolParameters) LoadParameters(params []string) error {
	switch toolParams.Action {
	case ActionTypeOktaApplicationCreate:
		toolParams.DataLocation = &(params[0])
		toolParams.TemplateName = &(params[1])

		length := len(params)
		if length > 2 {
			toolParams.TemplateLocation = &(params[2])
		} else {
			path, err := os.Getwd()

			if err != nil {
				return errors.New("Error finding current directory")
			}

			fullPath := path + OktaParamsDefaultTemplateLocation
			toolParams.TemplateLocation = &fullPath
		}

		if length > 3 {
			toolParams.TemplateLocation = &(params[3])
		}

		return nil
	case ActionTypeOktaApplicationList,
		ActionTypeOktaApplicationDisable,
		ActionTypeOktaApplicationEnable,
		ActionTypeOktaApplicationDelete:
		toolParams.DataLocation = &(params[0])
		return nil
	case ActionTypeOktaApplicationUpdateCert:
		toolParams.DataLocation = &(params[0])
		toolParams.CertLocation = &(params[1])
		return nil
	}

	return errors.New("Action " + toolParams.Action + " is not supported")
}
