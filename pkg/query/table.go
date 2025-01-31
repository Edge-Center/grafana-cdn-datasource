package query

import (
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func NewTableFrame(qm models.QueryModel, response *statistics.ResourceStatisticsTableResponse) []*data.Frame {
	var metrics = qm.Metrics
	var frames []*data.Frame

	if response == nil {
		return frames
	}

	for _, metric := range metrics {
		for _, tableData := range *response {
			labels := NewTableLabels(&tableData, metric, qm)
			name := GetMetricLabel(metric, models.PluginStrings)

			switch metric {
			case string(statistics.MetricUpstreamBytes):
				frames = append(frames, NewTableFrameForMetricBytes(name, tableData.Metrics.UpstreamBytes, labels, qm))
			case string(statistics.MetricSentBytes):
				frames = append(frames, NewTableFrameForMetricBytes(name, tableData.Metrics.SentBytes, labels, qm))
			case string(statistics.MetricShieldBytes):
				frames = append(frames, NewTableFrameForMetricBytes(name, tableData.Metrics.ShieldBytes, labels, qm))
			case string(statistics.MetricTotalBytes):
				frames = append(frames, NewTableFrameForMetricBytes(name, tableData.Metrics.TotalBytes, labels, qm))
			case string(statistics.MetricCdnBytes):
				frames = append(frames, NewTableFrameForMetricBytes(name, tableData.Metrics.CDNBytes, labels, qm))
			case string(statistics.MetricRequestTime):
				frames = append(frames, NewTableFrameForMetricTime(name, tableData.Metrics.RequestTime, labels, qm))
			case string(statistics.MetricOriginResponseTime):
				frames = append(frames, NewTableFrameForMetricTime(name, tableData.Metrics.OriginResponseTime, labels, qm))
			case string(statistics.MetricRequests):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.Requests, labels, qm))
			case string(statistics.MetricResponses2xx):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.Responses2xx, labels, qm))
			case string(statistics.MetricResponses3xx):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.Responses3xx, labels, qm))
			case string(statistics.MetricResponses4xx):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.Responses4xx, labels, qm))
			case string(statistics.MetricResponses5xx):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.Responses5xx, labels, qm))
			case string(statistics.MetricResponsesHit):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.ResponsesHit, labels, qm))
			case string(statistics.MetricResponsesMiss):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.ResponsesMiss, labels, qm))
			case string(statistics.MetricImageProcessed):
				frames = append(frames, NewTableFrameForMetricInt(name, tableData.Metrics.ImageProcessed, labels, qm))
			case string(statistics.MetricCacheHitTrafficRatio):
				frames = append(frames, NewTableFrameForMetricPercent(name, tableData.Metrics.CacheHitTrafficRatio, labels, qm))
			case string(statistics.MetricCacheHitRequestsRatio):
				frames = append(frames, NewTableFrameForMetricPercent(name, tableData.Metrics.CacheHitRequestsRatio, labels, qm))
			case string(statistics.MetricShieldTrafficRatio):
				frames = append(frames, NewTableFrameForMetricPercent(name, tableData.Metrics.ShieldTrafficRatio, labels, qm))
			case string(models.PluginMetricBandwidth):
				frames = append(frames, NewTableFrameForMetricBandwidth(name, tableData.Metrics.TotalBytes, labels, qm))
			}
		}
	}

	return frames
}

func NewTableFrameForMetricBytes(name string, value uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	decimals := uint16(2)

	valueField := data.NewField(TableValueFieldName, labels, []uint64{value}).SetConfig(&data.FieldConfig{
		Unit:              "decbytes",
		Decimals:          &decimals,
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, valueField)

	return frame
}

func NewTableFrameForMetricPercent(name string, value float64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	decimals := uint16(2)
	minValue := data.ConfFloat64(0)
	maxValue := data.ConfFloat64(1)

	valueField := data.NewField(TableValueFieldName, labels, []float64{clamp(value, 0, 1)}).SetConfig(&data.FieldConfig{
		Unit:              "percentunit",
		Decimals:          &decimals,
		Min:               &minValue,
		Max:               &maxValue,
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, valueField)

	return frame
}

func NewTableFrameForMetricInt(name string, value uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)

	valueField := data.NewField(TableValueFieldName, labels, []uint64{value}).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, valueField)

	return frame
}

func NewTableFrameForMetricTime(name string, value uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)

	valueField := data.NewField(TableValueFieldName, labels, []uint64{value}).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
		Unit:              "ms",
	})

	frame.Fields = append(frame.Fields, valueField)

	return frame
}

func NewTableFrameForMetricBandwidth(name string, value uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	decimals := uint16(2)
	period, err := statistics.GranularityToSeconds(qm.Granularity)
	if err != nil {
		return frame
	}

	bandwidth := float64(value*8) / float64(period)

	valueField := data.NewField(TableValueFieldName, labels, []float64{bandwidth}).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
		Unit:              "Bps",
		Decimals:          &decimals,
	})

	frame.Fields = append(frame.Fields, valueField)

	return frame
}
