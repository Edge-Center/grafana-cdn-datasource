import { useCallback } from 'react';
import { EditorProps } from './types';
import { SecureJsonData } from '../../types';

type OnChangeType = () => void;

export function useResetSecureOptions(props: EditorProps, propertyName: keyof SecureJsonData): OnChangeType {
  const { onOptionsChange, options } = props;

  return useCallback(() => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        [propertyName]: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        [propertyName]: '',
      },
    });
  }, [onOptionsChange, options, propertyName]);
}
