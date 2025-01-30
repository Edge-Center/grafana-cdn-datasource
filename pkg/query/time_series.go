package query

import (
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/Edge-Center/grafana-cdn-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"time"
)

func NewTimeSeriesFrame(qm models.QueryModel, response *statistics.ResourceStatisticsTimeSeriesResponse) []*data.Frame {
	var metrics = qm.Metrics
	var frames []*data.Frame

	for _, metric := range metrics {
		for _, timeSeriesData := range *response {
			labels := NewLabels(&timeSeriesData, metric, qm)

			switch metric {
			case string(statistics.MetricUpstreamBytes):
				frames = append(frames, NewFrameForMetricBytes(string(metric), timeSeriesData.Metrics.UpstreamBytes, labels, qm))
			case string(statistics.MetricSentBytes):
				frames = append(frames, NewFrameForMetricBytes(string(metric), timeSeriesData.Metrics.SentBytes, labels, qm))
			case string(statistics.MetricShieldBytes):
				frames = append(frames, NewFrameForMetricBytes(string(metric), timeSeriesData.Metrics.ShieldBytes, labels, qm))
			case string(statistics.MetricTotalBytes):
				frames = append(frames, NewFrameForMetricBytes(string(metric), timeSeriesData.Metrics.TotalBytes, labels, qm))
			case string(statistics.MetricCdnBytes):
				frames = append(frames, NewFrameForMetricBytes(string(metric), timeSeriesData.Metrics.CDNBytes, labels, qm))
			case string(statistics.MetricRequestTime):
				frames = append(frames, NewFrameForMetricTime(string(metric), timeSeriesData.Metrics.RequestTime, labels, qm))
			case string(statistics.MetricOriginResponseTime):
				frames = append(frames, NewFrameForMetricTime(string(metric), timeSeriesData.Metrics.OriginResponseTime, labels, qm))
			case string(statistics.MetricRequests):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.Requests, labels, qm))
			case string(statistics.MetricResponses2xx):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.Responses2xx, labels, qm))
			case string(statistics.MetricResponses3xx):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.Responses3xx, labels, qm))
			case string(statistics.MetricResponses4xx):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.Responses4xx, labels, qm))
			case string(statistics.MetricResponses5xx):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.Responses5xx, labels, qm))
			case string(statistics.MetricResponsesHit):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.ResponsesHit, labels, qm))
			case string(statistics.MetricResponsesMiss):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.ResponsesMiss, labels, qm))
			case string(statistics.MetricImageProcessed):
				frames = append(frames, NewFrameForMetricInt(string(metric), timeSeriesData.Metrics.ImageProcessed, labels, qm))
			case string(statistics.MetricCacheHitTrafficRatio):
				frames = append(frames, NewFrameForMetricPercent(string(metric), timeSeriesData.Metrics.CacheHitTrafficRatio, labels, qm))
			case string(statistics.MetricCacheHitRequestsRatio):
				frames = append(frames, NewFrameForMetricPercent(string(metric), timeSeriesData.Metrics.CacheHitRequestsRatio, labels, qm))
			case string(statistics.MetricShieldTrafficRatio):
				frames = append(frames, NewFrameForMetricPercent(string(metric), timeSeriesData.Metrics.ShieldTrafficRatio, labels, qm))
			case string(PluginMetricBandwidth):
				frames = append(frames, NewFrameForMetricBandwidth(string(metric), timeSeriesData.Metrics.TotalBytes, labels, qm))
			}
		}
	}

	return frames
}

func NewFrameForMetricBytes(name string, dataPoints [][]uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]uint64, 0, len(dataPoints))
	decimals := uint16(2)

	for _, point := range dataPoints {
		if len(point) < 2 {
			continue
		}
		timestamps = append(timestamps, time.Unix(int64(point[0]), 0))
		values = append(values, point[1])
	}

	timeField := data.NewField(TimeSeriesTimeFieldName, nil, timestamps)

	valueField := data.NewField(TimeSeriesValuesFieldName, labels, values).SetConfig(&data.FieldConfig{
		Unit:              "bytes",
		Decimals:          &decimals,
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}

func NewFrameForMetricPercent(name string, dataPoints [][]float64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]float64, 0, len(dataPoints))
	decimals := uint16(2)
	minValue := data.ConfFloat64(0)
	maxValue := data.ConfFloat64(1)

	for _, point := range dataPoints {
		if len(point) < 2 {
			continue
		}
		timestamps = append(timestamps, time.Unix(int64(point[0]), 0))
		values = append(values, clamp(point[1], 0, 1))
	}

	timeField := data.NewField(TimeSeriesTimeFieldName, nil, timestamps)

	valueField := data.NewField(TimeSeriesValuesFieldName, labels, values).SetConfig(&data.FieldConfig{
		Unit:              "percentunit",
		Decimals:          &decimals,
		Min:               &minValue,
		Max:               &maxValue,
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}

func NewFrameForMetricInt(name string, dataPoints [][]uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]uint64, 0, len(dataPoints))

	for _, point := range dataPoints {
		if len(point) < 2 {
			continue
		}
		timestamps = append(timestamps, time.Unix(int64(point[0]), 0))
		values = append(values, point[1])
	}

	timeField := data.NewField(TimeSeriesTimeFieldName, nil, timestamps)

	valueField := data.NewField(TimeSeriesValuesFieldName, labels, values).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}

func NewFrameForMetricTime(name string, dataPoints [][]uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]uint64, 0, len(dataPoints))

	for _, point := range dataPoints {
		if len(point) < 2 {
			continue
		}
		timestamps = append(timestamps, time.Unix(int64(point[0]), 0))
		values = append(values, point[1])
	}

	timeField := data.NewField(TimeSeriesTimeFieldName, nil, timestamps)

	valueField := data.NewField(TimeSeriesValuesFieldName, labels, values).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
		Unit:              "ms",
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}

func NewFrameForMetricBandwidth(name string, dataPoints [][]uint64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]float64, 0, len(dataPoints))
	decimals := uint16(2)
	period, err := statistics.GranularityToSeconds(qm.Granularity)
	if err != nil {
		return frame
	}

	for _, point := range dataPoints {
		if len(point) < 2 {
			continue
		}
		timestamps = append(timestamps, time.Unix(int64(point[0]), 0))
		values = append(values, float64(point[1]*8)/float64(period))
	}

	timeField := data.NewField(TimeSeriesTimeFieldName, nil, timestamps)

	valueField := data.NewField(TimeSeriesValuesFieldName, labels, values).SetConfig(&data.FieldConfig{
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
		Unit:              "Bps",
		Decimals:          &decimals,
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}
