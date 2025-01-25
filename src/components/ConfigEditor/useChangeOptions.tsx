import { ChangeEvent, useCallback } from 'react';
import type { DataSourceOptions } from 'types';
import { EditorProps } from "./types";


type OnChangeType = (event: ChangeEvent<HTMLInputElement>) => void;

export function useChangeOptions(props: EditorProps, propertyName: keyof DataSourceOptions): OnChangeType {
  const { onOptionsChange, options } = props;

  return useCallback(
    (event: ChangeEvent<HTMLInputElement>) => {
      onOptionsChange({
        ...options,
        jsonData: {
          ...options.jsonData,
          [propertyName]: event.target.value,
        },
      });
    },
    [onOptionsChange, options, propertyName]
  );
}
