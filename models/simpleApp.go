package models

type SimpleApp struct {
	Id     string
	Code   string
	Url    string
	Label  string
	Status string
}

func (app SimpleApp) FindStorageAppInList(list []SimpleApp) *SimpleApp {
	for _, listApp := range list {
		if app.Label == listApp.Label {
			listApp.Url = app.Url
			return &listApp
		}
	}

	return nil
}
