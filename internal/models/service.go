package models

type Service struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Versions      []string `json:"versions"`
	TotalVersions int      `json:"total_versions"`
}

type Version struct {
	Id        string `json:"id"`
	ServiceId string `json:"service_id"`
	Name      string `json:"name"`
}

type QueryParams struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sort   string `json:"sort"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}
