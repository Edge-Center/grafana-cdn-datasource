import { useAsync } from 'react-use';
import { MetricsResponse, StringsResponse } from '../../types';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';
import {getMetricStrings, resolveDesc, resolveLabel} from './utils';

type AsyncQueryMetricsState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryMetrics(datasource: DataSource): AsyncQueryMetricsState {
  const result = useAsync(async () => {
    const [metricsData, strings]: [MetricsResponse, StringsResponse] = await Promise.all([
      datasource.getAvailableMetrics(),
      datasource.getStrings(),
    ]);

    return metricsData.metrics.map(
      (value) =>
        ({
          label: resolveLabel(value, getMetricStrings(strings)),
          description: resolveDesc(value, getMetricStrings(strings)),
          value,
        } as SelectableValue<string>)
    );
  }, [datasource]);

  return {
    loading: result.loading,
    options: result.value ?? [],
    error: result.error,
  };
}
