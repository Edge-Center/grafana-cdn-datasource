import {InlineField, Select} from "@grafana/ui";
import React, {FC} from "react";
import { OnChangeType } from "../useChangeSelectableValue";
import { DataSource } from "../../../datasource";
import { useQueryMetrics } from "../useQueryMetrics";
import { SelectableValue } from "@grafana/data";

export interface Props {
    datasource: DataSource
    value: SelectableValue<string> | undefined
    onChange: OnChangeType
}
export const MetricPicker: FC<Props> = ({value, datasource, onChange}) => {
    const { loading, options, error } = useQueryMetrics(datasource);

    return (
        <InlineField label="Metric" labelWidth={14}>
            <Select
                inputId="editor-query-metric"
                options={options}
                onChange={onChange}
                isLoading={loading}
                disabled={!!error}
                value={value}
            />
        </InlineField>
    )
}