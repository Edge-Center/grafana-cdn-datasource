package models

import "github.com/Edge-Center/edgecentercdn-go/statistics"

type QueryModel struct {
	QueryType    string                  `json:"queryType"`
	Metrics      []statistics.Metric     `json:"metrics"`
	Regions      []statistics.Region     `json:"regions,omitempty"`
	Vhosts       []statistics.Vhost      `json:"vhosts,omitempty"`
	Clients      []statistics.ClientId   `json:"clients,omitempty"`
	Countries    []statistics.Country    `json:"countries,omitempty"`
	Resources    []statistics.ResourceId `json:"resources,omitempty"`
	GroupBy      []statistics.GroupBy    `json:"groupby,omitempty"`
	Granularity  statistics.Granularity  `json:"granularity,omitempty"`
	LegendFormat string                  `json:"legendFormat,omitempty"`
}
