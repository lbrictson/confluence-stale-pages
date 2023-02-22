package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ConfluenceAPI struct {
	BaseURL string
	APIKey  string
	Email   string
}

func (c *ConfluenceAPI) ListAllContent() ([]ListContentItem, error) {
	offset := 0
	var content []ListContentItem
	for {
		page, more, err := c.listContentPage(offset)
		if err != nil {
			return nil, err
		}
		content = append(content, page...)
		if !more {
			break
		}
		offset += 1000
	}
	return content, nil
}

// listContentPage will return a list of content items from the Confluence API using the offset provided,
// if there is another page true will be returned
func (c *ConfluenceAPI) listContentPage(offset int) ([]ListContentItem, bool, error) {
	url := c.BaseURL + "/wiki/rest/api/content?limit=1000&start=" + fmt.Sprintf("%d", offset)
	method := "GET"
	// Customize later if need be
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, false, err
	}
	req.SetBasicAuth(c.Email, c.APIKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}
	data := ListContentResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, false, err
	}
	more := false
	if data.Links.Next != "" {
		more = true
	}
	return data.Results, more, nil
}
