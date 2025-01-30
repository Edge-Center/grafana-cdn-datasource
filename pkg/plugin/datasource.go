package plugin

import (
	"context"
	"fmt"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/client"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/query"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

var (
	_ backend.QueryDataHandler      = (*Datasource)(nil)
	_ backend.CheckHealthHandler    = (*Datasource)(nil)
	_ instancemgmt.InstanceDisposer = (*Datasource)(nil)
)

func NewDatasource(_ context.Context, dis backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	settings, err := models.LoadPluginSettings(dis)
	if err != nil {
		return nil, err
	}

	return &Datasource{
		CallResourceHandler: NewResourceHandler(settings),
		settings:            settings,
	}, nil
}

type Datasource struct {
	backend.CallResourceHandler
	settings *models.PluginSettings
}

func (d *Datasource) Dispose() {
}

func (d *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()

	for _, q := range req.Queries {
		res := query.RunQuery(ctx, d.settings, q)

		response.Responses[q.RefID] = res
	}

	return response, nil
}

func (d *Datasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	pluginSettings, err := models.LoadPluginSettings(*req.PluginContext.DataSourceInstanceSettings)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: "Unable to load settings",
		}, nil
	}

	cdnClient, err := client.NewCdnServicePluginSettings(pluginSettings)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	username, err := cdnClient.Tools().Whoami(ctx)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: fmt.Sprintf("You successfully authenticated as %s", username),
	}, nil
}
