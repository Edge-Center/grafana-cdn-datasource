package models

import "github.com/Edge-Center/edgecentercdn-go/statistics"

type QueryModel struct {
	QueryType    string                  `json:"queryType"`
	Metrics      []string                `json:"metrics"`
	Regions      []statistics.Region     `json:"regions,omitempty"`
	Hosts        []statistics.Host       `json:"hosts,omitempty"`
	Clients      []statistics.ClientId   `json:"clients,omitempty"`
	Countries    []statistics.Country    `json:"countries,omitempty"`
	Resources    []statistics.ResourceId `json:"resources,omitempty"`
	GroupBy      []statistics.GroupBy    `json:"groupby,omitempty"`
	Granularity  statistics.Granularity  `json:"granularity,omitempty"`
	LegendFormat string                  `json:"legendFormat,omitempty"`
}
