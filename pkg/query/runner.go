package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/client"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func RunQuery(ctx context.Context, pluginSettings *models.PluginSettings, query backend.DataQuery) backend.DataResponse {
	var qm models.QueryModel
	if err := json.Unmarshal(query.JSON, &qm); err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("json unmarshal: %v", err))
	}

	cdnClient, err := client.NewCdnServicePluginSettings(pluginSettings)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("create cdn client: %v", err))
	}

	var frames []*data.Frame
	if qm.QueryType == models.QueryTypeTable {
		stats, _ := cdnClient.Statistics().GetTableData(ctx, createTableRequest(qm, query))
		frames = NewTableFrame(qm, stats)
	} else {
		stats, _ := cdnClient.Statistics().GetTimeSeriesData(ctx, createTimeSeriesRequest(qm, query))
		frames = NewTimeSeriesFrame(qm, stats)
	}

	for _, frame := range frames {
		frame.RefID = query.RefID
	}

	return backend.DataResponse{Frames: frames}
}

func createTimeSeriesRequest(qm models.QueryModel, query backend.DataQuery) *statistics.ResourceStatisticsTimeSeriesRequest {
	return &statistics.ResourceStatisticsTimeSeriesRequest{
		Metrics:     GetMetrics(qm.Metrics),
		Regions:     qm.Regions,
		GroupBy:     qm.GroupBy,
		Granularity: qm.Granularity,
		Hosts:       qm.Hosts,
		Resources:   qm.Resources,
		Countries:   qm.Countries,
		From:        query.TimeRange.From,
		To:          query.TimeRange.To,
	}
}

func createTableRequest(qm models.QueryModel, query backend.DataQuery) *statistics.ResourceStatisticsTableRequest {
	return &statistics.ResourceStatisticsTableRequest{
		Metrics:   GetMetrics(qm.Metrics),
		Regions:   qm.Regions,
		GroupBy:   qm.GroupBy,
		Hosts:     qm.Hosts,
		Resources: qm.Resources,
		Countries: qm.Countries,
		From:      query.TimeRange.From,
		To:        query.TimeRange.To,
	}
}
