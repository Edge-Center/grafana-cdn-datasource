package models

import "github.com/Edge-Center/edgecentercdn-go/statistics"

type PluginMetric string

const (
	PluginMetricBandwidth PluginMetric = "bandwidth"
)

var MetricsSuggestions = append(statistics.MetricsSuggestions, string(PluginMetricBandwidth))

var PluginMetricStrings = map[PluginMetric]statistics.StringDetails{
	PluginMetricBandwidth: {
		Label: "Bandwidth",
		Desc:  "It is calculated based on the sum of the traffic from the origin to the CDN servers or shielding, the traffic from shielding to the CDN servers, and the traffic from the CDN servers to the end users.",
	},
}

var QueryTypesStrings = map[QueryType]statistics.StringDetails{
	QueryTypeTimeSeries: {
		Label: "Time Series",
		Desc:  "Returns data as [Time, Value] pairs",
	},
	QueryTypeTable: {
		Label: "Table",
		Desc:  "Returns a single value",
	},
}

type PluginStringMappings struct {
	statistics.StringMappings
	PluginMetrics map[PluginMetric]statistics.StringDetails `json:"pluginMetrics"`
	QueryTypes    map[QueryType]statistics.StringDetails    `json:"queryTypes"`
}

var PluginStrings = PluginStringMappings{
	StringMappings: statistics.Strings,
	PluginMetrics:  PluginMetricStrings,
	QueryTypes:     QueryTypesStrings,
}
