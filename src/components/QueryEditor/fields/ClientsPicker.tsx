import React, { FC } from 'react';
import { InlineField, Input } from '@grafana/ui';
import { useChangeInput } from '../useChangeInput';

export interface Props {
  value: string;
  onChange: (value: string) => void;
}
export const ClientsPicker: FC<Props> = ({ value, onChange }) => {
  const handleChange = useChangeInput(onChange);

  return (
    <InlineField
      label="Clients"
      tooltip="Filter by client IDs. Use commas to separate multiple values or reference variables."
      labelWidth={14}
      interactive
    >
      <Input
        id="editor-query-clients"
        onChange={handleChange}
        value={value}
        width={40}
        placeholder={'Enter client IDs separated by commas'}
      />
    </InlineField>
  );
};
