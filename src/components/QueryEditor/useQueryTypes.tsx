import { useAsync } from 'react-use';
import type { SelectableValue } from '@grafana/data';
import type { DataSource } from '../../datasource';
import { QueryTypesResponse, StringsResponse } from "../../types";
import { resolveDesc, resolveLabel } from "./utils";

type AsyncQueryTypeState = {
  loading: boolean;
  queryTypes: Array<SelectableValue<string>>;
  error: Error | undefined;
};

export function useQueryTypes(datasource: DataSource): AsyncQueryTypeState {
  const result = useAsync(async () => {
    const [queryTypesData, strings]: [QueryTypesResponse, StringsResponse] = await Promise.all([
      datasource.getAvailableQueryTypes(),
      datasource.getStrings(),
    ]);

    return queryTypesData.queryTypes.map(
        (value) =>
            ({
              label: resolveLabel(value, strings.queryTypes),
              description: resolveDesc(value, strings.queryTypes),
              value,
            } as SelectableValue<string>)
    );
  }, [datasource]);

  return {
    loading: result.loading,
    queryTypes: result.value ?? [],
    error: result.error,
  };
}
