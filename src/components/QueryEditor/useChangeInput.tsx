import { ChangeEvent, useCallback } from 'react';

export type OnChangeType = (value: string) => void;

export const useChangeInput = (onChange: OnChangeType) => {
  return useCallback(
    (event: ChangeEvent<HTMLInputElement>) => {
      onChange(event.target.value);
    },
    [onChange]
  );
};
