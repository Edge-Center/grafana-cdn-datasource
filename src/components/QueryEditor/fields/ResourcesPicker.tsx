import React, { FC } from 'react';
import { InlineField, Input } from '@grafana/ui';
import { useChangeInput } from '../useChangeInput';

export interface Props {
  value: string;
  onChange: (value: string) => void;
}
export const ResourcesPicker: FC<Props> = ({ value, onChange }) => {
  const handleChange = useChangeInput(onChange);

  return (
    <InlineField
      label="Resources"
      tooltip="Filter by resource IDs. Use commas to separate multiple values or reference variables."
      labelWidth={14}
      interactive
    >
      <Input
        id="editor-query-resources"
        onChange={handleChange}
        value={value}
        width={40}
        placeholder="Enter resource IDs separated by commas"
      />
    </InlineField>
  );
};
