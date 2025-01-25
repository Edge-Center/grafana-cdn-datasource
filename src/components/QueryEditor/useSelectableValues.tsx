import { useMemo } from 'react';
import type { SelectableValue } from '@grafana/data';

export function useSelectableValues(value: Array<string> | string | undefined, label?: Array<string>): Array<SelectableValue<string>> | undefined {
  return useMemo(() => {
    if (!value || value.length === 0) {
      return;
    }

    if (Array.isArray(value)) {
      const labels = label && label.length >= value.length ? label : value;


      return value.map((item, index) => ({
        label: labels[index] ?? item,
        value: item,
      }));
    }

    return [{
        label: label ?? value,
        value: value,
      }
    ]
  }, [label, value]);
}
