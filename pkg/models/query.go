package models

import "github.com/Edge-Center/edgecentercdn-go/statistics"

type QueryType = string

const (
	QueryTypeTimeSeries = "timeSeries"
	QueryTypeTable      = "table"
)

type QueryModel struct {
	QueryType    QueryType               `json:"queryType"`
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
