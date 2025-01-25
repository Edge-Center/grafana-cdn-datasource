package client

import (
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go"
	"github.com/Edge-Center/edgecentercdn-go/edgecenter/provider"
	"github.com/edge-center/cdn-datasource/pkg/models"
	"net/http"
)

func NewCdnServicePluginSettings(pluginSettings *models.PluginSettings) (*edgecentercdn_go.Service, error) {
	apiUrl := pluginSettings.ApiUrl
	if apiUrl == "" {
		return nil, fmt.Errorf("API url is missing")
	}

	apiKey := pluginSettings.Secrets.ApiKey
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	var opts []provider.ClientOption

	if apiKey != "" {
		opts = append(opts, provider.WithSignerFunc(createSignerFunc(apiKey)))
	}

	client := provider.NewClient(apiUrl, opts...)

	service := edgecentercdn_go.NewService(client)

	return service, nil
}

func createSignerFunc(apiKey string) func(req *http.Request) error {
	return func(req *http.Request) error {
		for k, v := range provider.AuthenticatedHeaders(apiKey) {
			req.Header.Set(k, v)
		}
		return nil
	}
}
