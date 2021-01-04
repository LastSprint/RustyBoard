package jira

import (
	"github.com/LastSprint/JiraGoIssues/models"
	"github.com/LastSprint/JiraGoIssues/services"
)

type JiraServiceWrapper struct {
	*services.JiraIssueLoader
}

func (srv *JiraServiceWrapper) LoadIssues(jql string) (models.IssueSearchWrapperEntity, error) {
	return srv.JiraIssueLoader.LoadIssues(StringWrapper(jql))
}

type StringWrapper string

func (strw StringWrapper) GetUseOnlyAdditionalFields() bool {
	return false
}

// MakeJiraRequest явно вызывает конвертацию сущности в запрос.
func (strw StringWrapper) MakeJiraRequest() string {
	return string(strw)
}

// AdditionFields возвращает список полей, которые ожидается получить от запроса. Этот список дополняется к списку по-умолчанию.
func (strw StringWrapper) GetAdditionFields() []services.JiraField {
	return []services.JiraField{}
}
