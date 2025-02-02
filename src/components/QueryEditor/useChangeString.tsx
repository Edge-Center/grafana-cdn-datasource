import { useCallback } from 'react';
import type { Query } from '../../types';
import type { ChangeOptions, EditorProps } from './types';

type OnChangeType = (value: string) => void;

export function useChangeString(props: EditorProps, options: ChangeOptions<Query>): OnChangeType {
  const { onChange, onRunQuery, query } = props;
  const { propertyName, runQuery } = options;

  return useCallback(
    (value: string) => {
      onChange({
        ...query,
        [propertyName]: value,
      });

      if (runQuery) {
        onRunQuery();
      }
    },
    [onChange, onRunQuery, query, propertyName, runQuery]
  );
}
