import { useAsync } from 'react-use';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';

type AsyncQueryGroupsState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryGroupBy(datasource: DataSource): AsyncQueryGroupsState {
  const result = useAsync(async () => {
    const { groups } = await datasource.getAvailableGroups();

    return groups.map((group) => ({
      label: group,
      value: group,
    }));
  }, [datasource]);

  return {
    loading: result.loading,
    options: result.value ?? [],
    error: result.error,
  };
}
