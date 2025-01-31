package query

import (
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
)

func GetMetricLabel(metric string, mappings models.PluginStringMappings) string {
	if metricLabel, exists := mappings.Metrics[statistics.Metric(metric)]; exists {
		return metricLabel.Label
	}

	if metricLabel, exists := mappings.PluginMetrics[models.PluginMetric(metric)]; exists {
		return metricLabel.Label
	}

	return metric
}
