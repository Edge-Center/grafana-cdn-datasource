import { useAsync } from 'react-use';
import { GranularityResponse, StringsResponse } from '../../types';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';
import { resolveDesc, resolveLabel } from './utils';

type AsyncQueryGranularityState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryGranularity(datasource: DataSource): AsyncQueryGranularityState {
  const result = useAsync(async () => {
    const [granularityData, strings]: [GranularityResponse, StringsResponse] = await Promise.all([
      datasource.getAvailableGranularity(),
      datasource.getStrings(),
    ]);

    return granularityData.granularity.map(
      (value) =>
        ({
          label: resolveLabel(value, strings.granularity),
          description: resolveDesc(value, strings.granularity),
          value: value,
        } as SelectableValue<string>)
    );
  }, [datasource]);

  return {
    loading: result.loading,
    options: result.value ?? [],
    error: result.error,
  };
}
