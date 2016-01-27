package jira

import (
	"fmt"
	"strings"
)

type ChangeLog struct {
	StartAt    int       `json:"startAt,omitempty"`
	MaxResults int       `json:"maxResults,omitempty"`
	Total      int       `json:"total,omitempty"`
	Histories  Histories `json:"histories,omitempty"`
}

type Histories []*History

type History struct {
	Id      string  `json:"id,omitempty"`
	Author  *User   `json:"author,omitempty"`
	Created string  `json:"created,omitempty"`
	Items   []*Item `json:"items,omitempty"`
}

type Item struct {
	Field      string `json:"field,omitempty"`
	FieldType  string `json:"fieldtype,omitempty"`
	From       string `json:"from,omitempty"`
	FromString string `json:"fromString,omitempty"`
	To         string `json:"to,omitempty"`
	ToString   string `json:"toString,omitempty"`
}

func (history *History) ShowItems() string {
	var items []string
	if len(history.Items) > 0 {
		for _, item := range history.Items {
			items = append(items, fmt.Sprintf("%+v\n", *item))
		}
		return fmt.Sprintf("%v", strings.Join(items[:], "\n"))
	} else {
		return ""
	}
}

// Sorting methods for sorting histories by date DESC
func (slice Histories) Len() int {
	return len(slice)
}

func (slice Histories) Less(i, j int) bool {
	return slice[i].Created > slice[j].Created
}

func (slice Histories) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
