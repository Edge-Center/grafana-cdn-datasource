import { useAsync } from 'react-use';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';

type AsyncQueryGranularityState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryGranularity(datasource: DataSource): AsyncQueryGranularityState {
  const result = useAsync(async () => {
    const { granularity } = await datasource.getAvailableGranularity();

    return granularity.map((value) => ({
      label: value,
      value: value,
    }));
  }, [datasource]);

  return {
    loading: result.loading,
    options: result.value ?? [],
    error: result.error,
  };
}
