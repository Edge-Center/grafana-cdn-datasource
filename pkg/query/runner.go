package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/edge-center/cdn-datasource/pkg/client"
	"github.com/edge-center/cdn-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func RunQuery(ctx context.Context, pluginSettings *models.PluginSettings, query backend.DataQuery) backend.DataResponse {
	response := backend.DataResponse{}

	var qm models.QueryModel

	err := json.Unmarshal(query.JSON, &qm)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("json unmarshal: %v", err.Error()))
	}

	cdnClient, err := client.NewCdnServicePluginSettings(pluginSettings)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("create cdn client: %v", err.Error()))
	}

	req := &statistics.ResourceStatisticsTimeSeriesRequest{
		Metrics:     qm.Metrics,
		Regions:     qm.Regions,
		GroupBy:     qm.GroupBy,
		Granularity: qm.Granularity,
		VHosts:      qm.Vhosts,
		Resources:   qm.Resources,
		Countries:   qm.Countries,
		From:        query.TimeRange.From,
		To:          query.TimeRange.To,
	}

	log.DefaultLogger.Info(fmt.Sprintf("stats request: %+v", req.ToPath()))

	resp, err := cdnClient.Statistics().GetTimeSeriesData(ctx, req)

	if resp == nil {
		return response
	}

	log.DefaultLogger.Info(fmt.Sprintf("stats response: %+v", resp))

	frames := NewTimeSeriesFrame(query, qm, resp)

	if len(frames) == 0 {
		return response
	}

	for _, frame := range frames {
		frame.RefID = query.RefID
		response.Frames = append(response.Frames, frame)
	}

	return response
}
