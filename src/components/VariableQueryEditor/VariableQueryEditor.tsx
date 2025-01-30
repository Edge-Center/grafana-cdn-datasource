import React, {FC, useCallback} from 'react';
import { InlineField, Select } from '@grafana/ui';
import { Variable, VariableQuery } from '../../types';
import { useVariableType } from "./useVariableType";
import { SelectableValue } from "@grafana/data";

export interface Props {
  query: VariableQuery;
  onChange: (query: VariableQuery, definition: string) => void;
}

export const VariableQueryEditor: FC<Props> = ({ onChange, query }) => {
  const onSelectorChange = useCallback((value: SelectableValue<Variable>) => {
      onChange({selector: value}, value.label || '');
  }, [onChange]);

  const options = useVariableType();

  return (
    <InlineField label="Values for" labelWidth={16}>
      <Select
        width={16}
        maxVisibleValues={20}
        minMenuHeight={45}
        menuPlacement={'bottom'}
        onChange={onSelectorChange}
        value={query.selector}
        options={options}
      />
    </InlineField>
  );
};
