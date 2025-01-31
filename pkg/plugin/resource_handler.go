package plugin

import (
	"context"
	"encoding/json"
	"github.com/Edge-Center/edgecentercdn-go/resources"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/client"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
	"net/http"

	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
)

type contextKeyForPluginSettings string

const pluginSettingsContextKey contextKeyForPluginSettings = "plugin_settings"

func NewResourceHandler(pluginSettings *models.PluginSettings) backend.CallResourceHandler {
	mux := http.NewServeMux()
	mux.HandleFunc("/resources", withPluginSettings(handleResources, pluginSettings))
	mux.HandleFunc("/metrics", handleMetrics)
	mux.HandleFunc("/regions", handleRegions)
	mux.HandleFunc("/groups", handleGroups)
	mux.HandleFunc("/granularity", handleGranularity)
	mux.HandleFunc("/strings", handleStrings)
	mux.HandleFunc("/query-types", handleQueryTypes)

	return httpadapter.New(mux)
}

func withPluginSettings(handler http.HandlerFunc, pluginSettings *models.PluginSettings) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), pluginSettingsContextKey, pluginSettings)
		handler(w, r.WithContext(ctx))
	}
}

func handleResources(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()

	pluginSettings, ok := ctx.Value(pluginSettingsContextKey).(*models.PluginSettings)
	if !ok || pluginSettings == nil {
		http.Error(w, "pluginSettings not found in context", http.StatusInternalServerError)
		return
	}

	cdnClient, err := client.NewCdnServicePluginSettings(pluginSettings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := &resources.ListFilterRequest{
		Fields:  []string{"id", "cname", "client"},
		Deleted: true,
	}

	list, err := cdnClient.Resources().List(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out := map[string][]resources.Resource{
		"resources": list,
	}

	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	out := map[string][]string{
		"metrics": models.MetricsSuggestions,
	}

	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleRegions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	out := map[string][]string{
		"metrics": statistics.MetricsSuggestions,
	}

	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	out := map[string][]string{
		"groups": statistics.GroupBySuggestions,
	}

	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGranularity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	out := map[string][]string{
		"granularity": statistics.GranularitySuggestions,
	}
	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleStrings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	j, err := json.Marshal(models.PluginStrings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleQueryTypes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	out := map[string][]string{
		"queryTypes": []string{
			models.QueryTypeTimeSeries,
			models.QueryTypeTable,
		},
	}

	j, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
