import React, { FC } from 'react';
import { InlineField, Input } from '@grafana/ui';
import { useChangeInput } from '../useChangeInput';
import type { DataSource } from '../../../datasource';

export interface Props {
  datasource: DataSource;
  value: string;
  onChange: (value: string) => void;
}
export const HostsPicker: FC<Props> = ({ value, onChange }) => {
  const handleChange = useChangeInput(onChange);

  return (
    <InlineField
      label="Hosts"
      tooltip={`Filter by hosts. Use commas to separate multiple values or reference variables.`}
      labelWidth={14}
      interactive
    >
      <Input
        id="editor-query-hosts"
        onChange={handleChange}
        value={value}
        width={40}
        placeholder="Enter hosts separated by commas"
      />
    </InlineField>
  );
};
