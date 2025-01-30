import { useMemo } from 'react';
import { Strings } from '../../types';
import { resolveLabel } from './utils';
import type { SelectableValue } from '@grafana/data';

export function useSelectableValues(
  values: string[] | string | undefined,
  strings?: Strings
): Array<SelectableValue<string>> | undefined {
  return useMemo(() => {
    if (!values || values.length === 0) {
      return;
    }

    if (Array.isArray(values)) {
      return values.map((value, index) => ({
        label: resolveLabel(value, strings),
        value: value,
      }));
    }

    return [
      {
        label: resolveLabel(values, strings),
        value: values,
      },
    ] as Array<SelectableValue<string>>;
  }, [strings, values]);
}
