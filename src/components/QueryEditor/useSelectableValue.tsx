import { useMemo } from 'react';
import { Strings } from '../../types';
import { resolveLabel } from './utils';
import type { SelectableValue } from '@grafana/data';

export function useSelectableValue(value: string | undefined, strings?: Strings): SelectableValue<string> | undefined {
  return useMemo(() => {
    if (!value) {
      return;
    }

    return {
      label: resolveLabel(value, strings),
      value: value,
    } as SelectableValue<string>;
  }, [strings, value]);
}
