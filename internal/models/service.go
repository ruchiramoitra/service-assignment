package models

type Service struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type QueryParams struct {
	Name   string `json:"name"`
	Sort   string `json:"sort"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}
