import { InlineField, MultiSelect } from '@grafana/ui';
import React, { FC } from 'react';
import { OnChangeType } from '../useChangeSelectableValues';
import { DataSource } from '../../../datasource';
import { useQueryGroupBy } from '../useQueryGroupBy';
import { SelectableValue } from '@grafana/data';

export interface Props {
  datasource: DataSource;
  value: Array<SelectableValue<string>> | undefined;
  onChange: OnChangeType;
}

export const GroupByPicker: FC<Props> = ({ value, datasource, onChange }) => {
  const { loading, options, error } = useQueryGroupBy(datasource);

  return (
    <InlineField label="Group by" tooltip={`Fields used for grouping`} labelWidth={15} interactive>
      <MultiSelect
        inputId="editor-query-groupby"
        width={40}
        isMulti={true}
        options={options}
        onChange={onChange}
        isLoading={loading}
        disabled={!!error}
        value={value}
      />
    </InlineField>
  );
};
