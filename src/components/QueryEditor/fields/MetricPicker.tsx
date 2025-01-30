import { InlineField, MultiSelect } from '@grafana/ui';
import React, { FC } from 'react';
import { OnChangeType } from '../useChangeSelectableValues';
import { DataSource } from '../../../datasource';
import { useQueryMetrics } from '../useQueryMetrics';
import { SelectableValue } from '@grafana/data';

export interface Props {
  datasource: DataSource;
  value: Array<SelectableValue<string>> | undefined;
  onChange: OnChangeType;
}
export const MetricPicker: FC<Props> = ({ value, datasource, onChange }) => {
  const { loading, options, error } = useQueryMetrics(datasource);

  return (
    <InlineField label="Metric" labelWidth={15} required>
      <MultiSelect
        width={40}
        isMulti={true}
        inputId="editor-query-metric"
        options={options}
        minMenuHeight={500}
        onChange={onChange}
        isLoading={loading}
        disabled={!!error}
        value={value}
      />
    </InlineField>
  );
};
