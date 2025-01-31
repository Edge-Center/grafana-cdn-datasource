import React, {FC} from "react";
import { InlineField, Select } from "@grafana/ui";
import { SelectableValue } from "@grafana/data";
import { useQueryTypes } from "../useQueryTypes";
import { DataSource } from "../../../datasource";
import { OnChangeType } from "../useChangeSelectableValue";

export interface Props {
    datasource: DataSource;
    value: SelectableValue<string> | undefined;
    onChange: OnChangeType;
}

export const QueryTypePicker: FC<Props> = ({datasource, value, onChange}) => {
    const { loading, queryTypes, error } = useQueryTypes(datasource);

    return (
        <InlineField label="Query type" labelWidth={15}>
            <Select
                inputId="editor-query-type"
                width={40}
                options={queryTypes}
                onChange={onChange}
                isLoading={loading}
                disabled={!!error}
                value={value}
            />
        </InlineField>
    )
}
