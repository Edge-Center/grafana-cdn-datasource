import { InlineField, Select } from '@grafana/ui';
import React, { FC } from 'react';
import { OnChangeType } from '../useChangeSelectableValue';
import { DataSource } from '../../../datasource';
import { useQueryGranularity } from '../useQueryGranularity';
import { SelectableValue } from '@grafana/data';

export interface Props {
  datasource: DataSource;
  value: SelectableValue<string> | undefined;
  onChange: OnChangeType;
}
export const GranularityPicker: FC<Props> = ({ value, datasource, onChange }) => {
  const { loading, options, error } = useQueryGranularity(datasource);

  return (
    <InlineField label="Granularity" tooltip="Time series granularity" labelWidth={15} interactive required>
      <Select
        inputId="editor-query-granularity"
        width={40}
        options={options}
        onChange={onChange}
        isLoading={loading}
        disabled={!!error}
        value={value}
      />
    </InlineField>
  );
};
