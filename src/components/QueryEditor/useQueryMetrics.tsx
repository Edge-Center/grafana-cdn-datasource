import { useAsync } from 'react-use';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';

type AsyncQueryMetricsState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryMetrics(datasource: DataSource): AsyncQueryMetricsState {
  const result = useAsync(async () => {
    const { metrics } = await datasource.getAvailableMetrics();

    return metrics.map((metric) => ({
      label: metric,
      value: metric,
    }));
  }, [datasource]);

  return {
    loading: result.loading,
    options: result.value ?? [],
    error: result.error,
  };
}
