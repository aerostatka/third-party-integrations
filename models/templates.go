package models

type Templates struct {
	Templates map[string]Template `json:"templates"`
}

type Template struct {
	Name        string           `json:"name,omitempty"`
	AppTemplate *OktaApplication `json:"app,omitempty"`
}

type OktaApplication struct {
	Credentials interface{}              `json:"credentials,omitempty"`
	Visibility  interface{}              `json:"visibility,omitempty"`
	Features    []string                 `json:"features,omitempty"`
	Id          string                   `json:"id,omitempty"`
	Label       string                   `json:"label,omitempty"`
	Name        string                   `json:"name,omitempty"`
	SignOnMode  string                   `json:"signOnMode,omitempty"`
	Settings    *OktaApplicationSettings `json:"settings,omitempty"`
}

type OktaApplicationSettings struct {
	App    *OktaApplicationSettingsApp    `json:"app,omitempty"`
	SignOn *OktaApplicationSettingsSignOn `json:"signOn,omitempty"`
}

type OktaApplicationSettingsApp map[string]interface{}

type OktaApplicationSettingsSignOn struct {
	AllowMultipleAcsEndpoints bool                                      `json:"allowMultipleAcsEndpoints,omitempty"`
	AcsEndpoints              []string                                  `json:"acsEndpoints,omitempty"`
	AttributeStatements       []OktaApplicationSettingsSignOnAttributes `json:"attributeStatements,omitempty"`
	Audience                  string                                    `json:"audience,omitempty"`
	AuthnContextClassRef      string                                    `json:"authnContextClassRef,omitempty"`
	DefaultRelayState         string                                    `json:"defaultRelayState,omitempty"`
	Destination               string                                    `json:"destination,omitempty"`
	DigestAlgorithm           string                                    `json:"digestAlgorithm,omitempty"`
	HonorForceAuthn           bool                                      `json:"honorForceAuthn,omitempty"`
	IdpIssuer                 string                                    `json:"idpIssuer,omitempty"`
	Recipient                 string                                    `json:"recipient,omitempty"`
	RequestCompressed         bool                                      `json:"requestCompressed,omitempty"`
	ResponseSigned            bool                                      `json:"responseSigned,omitempty"`
	SignatureAlgorithm        string                                    `json:"signatureAlgorithm,omitempty"`
	SpCertificate             string                                    `json:"spCertificate,omitempty"`
	Slo                       *OktaApplicationSettingsSignOnSlo         `json:"slo,omitempty"`
	SsoAcsUrl                 string                                    `json:"ssoAcsUrl,omitempty"`
	SubjectNameIdFormat       string                                    `json:"subjectNameIdFormat,omitempty"`
	SubjectNameIdTemplate     string                                    `json:"subjectNameIdTemplate,omitempty"`
}

type OktaApplicationSettingsSignOnAttributes struct {
	Name      string   `json:"name,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	Type      string   `json:"type,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type OktaApplicationSettingsSignOnSlo struct {
	Enabled   bool   `json:"enabled,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
}

func (app *OktaApplication) IsApplicationInstance() bool {
	return true
}
