package pkg

type ListContentItem struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Status              string `json:"status"`
	Title               string `json:"title"`
	MacroRenderedOutput struct {
	} `json:"macroRenderedOutput"`
	Extensions struct {
		Position int `json:"position"`
	} `json:"extensions"`
	Expandable struct {
		Container           string `json:"container"`
		Metadata            string `json:"metadata"`
		Restrictions        string `json:"restrictions"`
		History             string `json:"history"`
		Body                string `json:"body"`
		Version             string `json:"version"`
		Descendants         string `json:"descendants"`
		Space               string `json:"space"`
		ChildTypes          string `json:"childTypes"`
		SchedulePublishInfo string `json:"schedulePublishInfo"`
		Operations          string `json:"operations"`
		SchedulePublishDate string `json:"schedulePublishDate"`
		Children            string `json:"children"`
		Ancestors           string `json:"ancestors"`
	} `json:"_expandable"`
	Links struct {
		Self   string `json:"self"`
		Tinyui string `json:"tinyui"`
		Editui string `json:"editui"`
		Webui  string `json:"webui"`
	} `json:"_links"`
}

type ListContentResponse struct {
	Results []ListContentItem `json:"results"`
	Start   int               `json:"start"`
	Limit   int               `json:"limit"`
	Size    int               `json:"size"`
	Links   struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
		Self    string `json:"self"`
	} `json:"_links"`
}
