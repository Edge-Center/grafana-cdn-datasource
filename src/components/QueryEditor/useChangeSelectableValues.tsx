import { useCallback } from 'react';
import type { SelectableValue } from '@grafana/data';
import type { Query } from '../../types';
import type { ChangeOptions, EditorProps } from './types';

export type OnChangeType = (value: Array<SelectableValue<string>>) => void;

export function useChangeSelectableValues(props: EditorProps, options: ChangeOptions<Query>): OnChangeType {
    const { onChange, onRunQuery, query } = props;
    const { propertyName, runQuery } = options;

    return useCallback(
        (selectable: Array<SelectableValue<string>>) => {
            console.log(selectable);
            onChange({
                ...query,
                [propertyName]: selectable.map(item => item.value),
            });

            if (runQuery) {
                onRunQuery();
            }
        },
        [onChange, onRunQuery, query, propertyName, runQuery]
    );
}
