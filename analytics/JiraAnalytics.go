package analytics

import (
	"RustyBoard/jira"
	"RustyBoard/persistence"
)

type JiraAnalyzer interface {
	Analyze(jql []string, name string) (*jira.ProjectData, []error)
}

type JiraAnalytics struct {
	JiraAnalyzer
}

func (an *JiraAnalytics) Run(data []persistence.StoredItem) ([]jira.ProjectData, []error) {

	resultData := []jira.ProjectData{}
	resultErrs := []error{}

	dataChans := make([]chan *jira.ProjectData, len(data))
	errChans := make([]chan []error, len(data))

	for i, item := range data {
		dc := make(chan *jira.ProjectData)
		ec := make(chan []error)

		dataChans[i] = dc
		errChans[i] = ec

		go an.runForOneProject(item, dc, ec)
	}

	for i := 0; i < len(dataChans); i++ {
		dt := <-dataChans[i]
		et := <-errChans[i]

		for _, err := range et {
			resultErrs = append(resultErrs, err)
		}

		if dt != nil {
			dt.ImageUrl = data[i].ImageUrl
			resultData = append(resultData, *dt)
		}
	}

	return resultData, resultErrs
}

func (an *JiraAnalytics) runForOneProject(item persistence.StoredItem, dataChan chan *jira.ProjectData, errChan chan []error) {

	projectInfo, errs := an.JiraAnalyzer.Analyze(item.JiraQueries, item.ProjectName)

	dataChan <- projectInfo
	errChan <- errs

	close(dataChan)
	close(errChan)
}

func (an *JiraAnalytics) singleRun(data []persistence.StoredItem) ([]jira.ProjectData, []error) {
	resultData := []jira.ProjectData{}
	resultErrs := []error{}

	for _, item := range data {
		projectInfo, errs := an.JiraAnalyzer.Analyze(item.JiraQueries, item.ProjectName)
		for _, err := range errs {
			resultErrs = append(resultErrs, err)
		}

		if projectInfo != nil {
			resultData = append(resultData, *projectInfo)
		}
	}

	return resultData, resultErrs
}
