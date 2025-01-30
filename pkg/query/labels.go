package query

import (
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
)

func NewLabels(data *statistics.ResourceStatisticsTimeSeriesData, metric string, qm models.QueryModel) map[string]string {
	labels := make(map[string]string)

	for _, group := range qm.GroupBy {
		switch group {
		case statistics.GroupByResource:
			if data.Resource != nil {
				labels[string(group)] = fmt.Sprintf("%.0f", *data.Resource)
			}
		case statistics.GroupByRegion:
			if data.Region != nil {
				labels[string(group)] = *data.Region
			}
		case statistics.GroupByHost:
			if data.Host != nil {
				labels[string(group)] = *data.Host
			}
		case statistics.GroupByClient:
			if data.Client != nil {
				labels[string(group)] = fmt.Sprintf("%.0f", *data.Client)
			}
		case statistics.GroupByClientRegion:
			if data.ClientRegion != nil {
				labels[string(group)] = *data.ClientRegion
			}
		case statistics.GroupByCountry:
			if data.Country != nil {
				labels[string(group)] = *data.Country
			}
		}
	}

	labels["metric"] = metric

	return labels
}
