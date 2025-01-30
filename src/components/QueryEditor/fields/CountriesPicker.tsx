import React, { FC } from 'react';
import { InlineField, Input } from '@grafana/ui';
import { useChangeInput } from '../useChangeInput';

export interface Props {
  value: string;
  onChange: (value: string) => void;
}
export const CountriesPicker: FC<Props> = ({ value, onChange }) => {
  const handleChange = useChangeInput(onChange);

  return (
    <InlineField
      label="Countries"
      tooltip="Filter by countries. Use commas to separate multiple values or reference variables."
      labelWidth={14}
      interactive
    >
      <Input
        id="editor-query-countries"
        onChange={handleChange}
        value={value}
        placeholder="Enter countries separated by commas"
        width={40}
      />
    </InlineField>
  );
};
