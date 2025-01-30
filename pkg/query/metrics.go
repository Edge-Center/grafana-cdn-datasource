package query

import (
	"github.com/Edge-Center/edgecentercdn-go/statistics"
)

type PluginMetric string

const (
	PluginMetricBandwidth PluginMetric = "bandwidth"
)

var MetricsSuggestions = append(statistics.MetricsSuggestions, string(PluginMetricBandwidth))

var PluginMetricMapping = map[PluginMetric]statistics.Metric{
	PluginMetricBandwidth: statistics.MetricTotalBytes,
}

var PluginMetricStrings = map[PluginMetric]statistics.StringDetails{
	PluginMetricBandwidth: {
		Label: "Bandwidth",
		Desc:  "It is calculated based on the sum of the traffic from the origin to the CDN servers or shielding, the traffic from shielding to the CDN servers, and the traffic from the CDN servers to the end users.",
	},
}

type PluginStringMappings struct {
	statistics.StringMappings
	PluginMetrics map[PluginMetric]statistics.StringDetails `json:"pluginMetrics"`
}

var PluginStrings = PluginStringMappings{
	StringMappings: statistics.Strings,
	PluginMetrics:  PluginMetricStrings,
}

func GetMetrics(metrics []string) []statistics.Metric {
	var result []statistics.Metric

	for _, metricStr := range metrics {
		if metric, exists := PluginMetricMapping[PluginMetric(metricStr)]; exists {
			result = append(result, metric)
		} else {
			result = append(result, statistics.Metric(metricStr))
		}
	}
	return result
}
