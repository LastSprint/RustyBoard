package jira

import (
	"github.com/LastSprint/JiraGoIssues/models"
	"log"
	"strconv"
)

type IssueLoader interface {
	LoadIssues(jql string) (models.IssueSearchWrapperEntity, error)
}

/// Analyzer leads issues form Jira and create aggregations -- ProjectData
type Analyzer struct {
	IssueLoader
}

/// Analyze make request for each JQL
// Then merge results from each JQL request
// BUT
// it will return duplicated tasks by issue.key
func (a *Analyzer) Analyze(jql []string, name string) (*ProjectData, []error) {
	entities, err := a.loadAll(jql)

	if len(entities) == 0 {
		return nil, err
	}

	return &ProjectData{
		Name:               name,
		WhoWorks:           collectByUser(entities),
		WholeWorkAnalytics: makeWorkAnalytics(entities),
	}, err
}

func (a *Analyzer) loadAll(jqls []string) ([]models.IssueEntity, []error) {
	existed := map[string]bool{}
	result := []models.IssueEntity{}
	errs := []error{}
	for _, jql := range jqls {
		val, err := a.IssueLoader.LoadIssues(jql)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		for _, item := range val.Issues {
			if _, ok := existed[item.Key]; !ok {
				existed[item.Key] = true
				result = append(result, item)
			}
		}
	}

	return result, errs
}

func collectByUser(issues []models.IssueEntity) []PerUserAnalytics {
	groupedByUser := map[string][]models.IssueEntity{}

	for _, item := range issues {
		if len(groupedByUser[item.Fields.Assignee.Name]) == 0 {
			groupedByUser[item.Fields.Assignee.Name] = []models.IssueEntity{}
		}

		groupedByUser[item.Fields.Assignee.Name] = append(groupedByUser[item.Fields.Assignee.Name], item)
	}

	result := []PerUserAnalytics{}

	for user, issue := range groupedByUser {
		item := PerUserAnalytics{
			User: User{
				Name: user,
			},
			WorkAnalytics: makeWorkAnalytics(issue),
		}

		result = append(result, item)
	}

	return result
}

func makeWorkAnalytics(issues []models.IssueEntity) WorkAnalytics {

	result := WorkAnalytics{}

	for _, item := range issues {

		switch item.Fields.Issuetype.ID {
		case models.IssueTypeTask:
			result.TaskSpent += item.Fields.TimeSpend
		case models.IssueTypeBug:
			result.BugSpent += item.Fields.TimeSpend
		case models.IssueTypeServiceTask:
			result.ServiceSpent += item.Fields.TimeSpend
		case models.IssueTypeTestExecuted:
			result.TestSpent += item.Fields.TimeSpend
		}

		switch item.Fields.Status.ID {
		case models.DoneID, models.FeedbackID, models.ApprovedID:
			result.Done += 1
		default:
			result.ToDo += 1
		}

		if item.Fields.Issuetype.ID != models.IssueTypeEpic {
			result.WholeSpent += item.Fields.TimeSpend
		}
	}

	result.WorkLog = calculateAverageSpendPerWeek(issues, "")

	return result
}

func calculateAverageSpendPerWeek(issues []models.IssueEntity, userName string) []Week {

	// TODO: - Filter by user

	wholeLog := map[week][]models.HistoryEntity{}

	for _, item := range issues {
		grouped := groupLogByWeek(item.Changelog.Histories)

		for key, value := range grouped {

			if len(wholeLog[key]) == 0 {
				wholeLog[key] = []models.HistoryEntity{}
			}

			for _, v := range value {
				wholeLog[key] = append(wholeLog[key], v)
			}
		}
	}

	weeks := []Week{}

	for week, history := range wholeLog {

		newWeek := Week{
			Year:      week.year,
			Week:      week.week,
			TimeSpent: 0,
		}

		for _, h := range history {

			for _, hi := range h.HistoryItems {
				if hi.FieldType == models.TimeSpent {
					from, err := strconv.Atoi(hi.From)
					to, err := strconv.Atoi(hi.To)

					if err != nil {
						log.Println("[ERROR] while parsing string to int", hi.From, err.Error())
						continue
					}

					newWeek.TimeSpent += to - from
				}
			}
		}
		weeks = append(weeks, newWeek)
	}

	return weeks
}

type week struct {
	year int
	week int
}

func groupLogByWeek(history []models.HistoryEntity) map[week][]models.HistoryEntity {
	result := map[week][]models.HistoryEntity{}

	for _, item := range history {

		y, w := item.CreatedDateTime.Time.ISOWeek()

		newWeek := week{y, w}

		if len(result[newWeek]) == 0 {
			result[newWeek] = []models.HistoryEntity{}
		}

		result[newWeek] = append(result[newWeek], item)
	}

	return result
}
