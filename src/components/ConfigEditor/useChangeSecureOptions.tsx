import { ChangeEvent, useCallback } from 'react';
import type { SecureJsonData } from 'types';
import type { EditorProps } from './types';

type OnChangeType = (event: ChangeEvent<HTMLInputElement>) => void;

export function useChangeSecureOptions(props: EditorProps, propertyName: keyof SecureJsonData): OnChangeType {
  const { onOptionsChange, options } = props;

  return useCallback(
    (event: ChangeEvent<HTMLInputElement>) => {
      onOptionsChange({
        ...options,
        secureJsonData: {
          ...options.secureJsonData,
          [propertyName]: event.target.value,
        },
      });
    },
    [onOptionsChange, options, propertyName]
  );
}
