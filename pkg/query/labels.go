package query

import (
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go/statistics"
)

func generateLabels(data *statistics.ResourceStatisticsTimeSeriesData, groupBy []statistics.GroupBy, metric string) map[string]string {
	labels := make(map[string]string)

	for _, group := range groupBy {
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

	labels["metric"] = string(metric)

	return labels
}
