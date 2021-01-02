package models

import "RustyBoard/jira"

type GetAllProjectResponse struct {
	Errors []error `json:"errors"`
	Data []jira.ProjectData `json:"data"`
}