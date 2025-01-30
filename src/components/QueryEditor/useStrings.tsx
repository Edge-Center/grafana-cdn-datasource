import { useAsync } from 'react-use';
import type { DataSource } from '../../datasource';

export function useStrings(datasource: DataSource) {
  return useAsync(() => datasource.getStrings(), [datasource]);
}
