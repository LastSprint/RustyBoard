package jira

type Week struct {
	Year int `json:"year"`
	// Week ISO week number
	Week int `json:"week"`
	// TimeSpent is what time was spent for this week
	TimeSpent int `json:"timeSpent"`
}

type WorkAnalytics struct {
	// TaskSpent is time spent on tasks in seconds
	TaskSpent int `json:"taskSpent"`
	// BugSpent is time spent on bugs in seconds
	BugSpent int `json:"bugSpent"`
	// ServiceSpent is time spent on service tasks in seconds
	ServiceSpent int `json:"serviceSpent"`
	// TestSpent is time spent on test execution tasks in second
	TestSpent int `json:"testSpent"`

	// WholeSpent what time was spent entirely
	WholeSpent int `json:"wholeSpent"`
	// WholeEstimated sum of issues' estimate
	WholeEstimated int `json:"wholeEstimated"`
	// WorkLog is array of week where each week associated with time was spent in this week
	WorkLog []Week `json:"workLog"`

	// Done is number of done tasks
	Done int `json:"done"`
	// ToDo is number of tasks that haven't been done
	ToDo int `json:"toDo"`
}

type User struct {
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

type PerUserAnalytics struct {
	User          `json:"user"`
	WorkAnalytics `json:"workAnalytics"`
}

// ProjectData describes analytics about this project
type ProjectData struct {
	Name               string             `json:"name"`
	WhoWorks           []PerUserAnalytics `json:"whoWorks"`
	WholeWorkAnalytics WorkAnalytics      `json:"wholeWorkAnalytics"`
	ImageUrl           string             `json:"imageUrl"`
}
