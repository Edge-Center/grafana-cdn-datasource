import { useAsync } from 'react-use';
import { GroupsResponse, StringsResponse } from '../../types';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';
import { resolveDesc, resolveLabel } from './utils';

type AsyncQueryGroupsState = {
  loading: boolean;
  options: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryGroupBy(datasource: DataSource): AsyncQueryGroupsState {
  const result = useAsync(async () => {
    const [groupsData, strings]: [GroupsResponse, StringsResponse] = await Promise.all([
      datasource.getAvailableGroups(),
      datasource.getStrings(),
    ]);

    return groupsData.groups.map(
      (value) =>
        ({
          label: resolveLabel(value, strings.groupBy),
          description: resolveDesc(value, strings.groupBy),
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
