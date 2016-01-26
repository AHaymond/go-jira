/*
   Copyright (C) 2016 AHaymond

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program. If not, see {http://www.gnu.org/licenses/}.
*/

package jira

import (
	"fmt"
	"strings"
)

type ChangeLog struct {
	StartAt    int       `json:"startAt,omitempty"`
	MaxResults int       `json:"maxResults,omitempty"`
	Total      int       `json:"total,omitempty"`
	Histories  []History `json:"histories,omitempty"`
}

type History struct {
	Id      string `json:"id,omitempty"`
	Author  *User  `json:"author,omitempty"`
	Created string `json:"created,omitempty"`
	Items   []Item `json:"items,omitempty"`
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
			items = append(items, fmt.Sprintf("%+v\n", item))
		}
		return fmt.Sprintf("%v", strings.Join(items[:], "\n"))
	} else {
		return ""
	}
}
