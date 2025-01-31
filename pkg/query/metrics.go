package query

import (
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
)

var PluginMetricMapping = map[models.PluginMetric]statistics.Metric{
	models.PluginMetricBandwidth: statistics.MetricTotalBytes,
}

func GetMetrics(metrics []string) []statistics.Metric {
	var result []statistics.Metric

	for _, metricStr := range metrics {
		if metric, exists := PluginMetricMapping[models.PluginMetric(metricStr)]; exists {
			result = append(result, metric)
		} else {
			result = append(result, statistics.Metric(metricStr))
		}
	}
	return result
}
