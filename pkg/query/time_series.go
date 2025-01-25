package query

import (
	"fmt"
	"github.com/Edge-Center/edgecentercdn-go/statistics"
	"github.com/edge-center/cdn-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"time"
)

func NewTimeSeriesFrame(query backend.DataQuery, queryModel models.QueryModel, response *statistics.ResourceStatisticsTimeSeriesResponse) []*data.Frame {
	var metrics = queryModel.Metrics
	var groupBy = queryModel.GroupBy
	var frames []*data.Frame

	for _, metric := range metrics {
		for _, resourceData := range *response {
			labels := generateLabels(&resourceData, groupBy, metric)

			switch metric {
			case statistics.MetricUpstreamBytes:
				frames = append(frames, NewFrameForMetricBytes(string(metric), resourceData.Metrics.UpstreamBytes, labels, queryModel))
			case statistics.MetricSentBytes:
				frames = append(frames, NewFrameForMetricBytes(string(metric), resourceData.Metrics.SentBytes, labels, queryModel))
			case statistics.MetricShieldBytes:
				frames = append(frames, NewFrameForMetricBytes(string(metric), resourceData.Metrics.ShieldBytes, labels, queryModel))
			case statistics.MetricTotalBytes:
				frames = append(frames, NewFrameForMetricBytes(string(metric), resourceData.Metrics.TotalBytes, labels, queryModel))
			case statistics.MetricCdnBytes:
				frames = append(frames, NewFrameForMetricBytes(string(metric), resourceData.Metrics.CDNBytes, labels, queryModel))
			case statistics.MetricRequestTime:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.RequestTime, labels, queryModel))
			case statistics.MetricRequests:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.Requests, labels, queryModel))
			case statistics.MetricResponses2xx:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.Responses2xx, labels, queryModel))
			case statistics.MetricResponses3xx:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.Responses3xx, labels, queryModel))
			case statistics.MetricResponses4xx:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.Responses4xx, labels, queryModel))
			case statistics.MetricResponses5xx:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.Responses5xx, labels, queryModel))
			case statistics.MetricResponsesHit:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.ResponsesHit, labels, queryModel))
			case statistics.MetricResponsesMiss:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.ResponsesMiss, labels, queryModel))
			case statistics.MetricImageProcessed:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.ImageProcessed, labels, queryModel))
			case statistics.MetricOriginResponseTime:
				frames = append(frames, NewFrameForMetricInt(string(metric), resourceData.Metrics.OriginResponseTime, labels, queryModel))
			case statistics.MetricCacheHitTrafficRatio:
				frames = append(frames, NewFrameForMetricPercent(string(metric), resourceData.Metrics.CacheHitTrafficRatio, labels, queryModel))
			case statistics.MetricCacheHitRequestsRatio:
				frames = append(frames, NewFrameForMetricPercent(string(metric), resourceData.Metrics.CacheHitRequestsRatio, labels, queryModel))
			case statistics.MetricShieldTrafficRatio:
				frames = append(frames, NewFrameForMetricPercent(string(metric), resourceData.Metrics.ShieldTrafficRatio, labels, queryModel))
			}
		}
	}

	return frames
}

func generateLabels(data *statistics.ResourceStatisticsTimeSeriesData, groupBy []statistics.GroupBy, metric statistics.Metric) map[string]string {
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
		case statistics.GroupByVhost:
			if data.Vhost != nil {
				labels[string(group)] = *data.Vhost
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

func NewFrameForMetricFloat(name string, dataPoints [][]float64, labels map[string]string, qm models.QueryModel) *data.Frame {
	frame := data.NewFrame(name)
	timestamps := make([]time.Time, 0, len(dataPoints))
	values := make([]float64, 0, len(dataPoints))
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
		Decimals:          &decimals,
		DisplayNameFromDS: renderTemplate(qm.LegendFormat, labels),
	})

	frame.Fields = append(frame.Fields, timeField, valueField)

	return frame
}
