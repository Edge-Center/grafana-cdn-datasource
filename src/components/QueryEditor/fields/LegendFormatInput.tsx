import React, { FC } from "react";
import { InlineField, Input } from "@grafana/ui";
import { useChangeInput } from "../useChangeInput";

export interface Props {
    value: string;
    onChange: (value: string) => void
}
export const LegendFormatInput: FC<Props> = ({value, onChange}) => {
    const handleChange = useChangeInput(onChange);

    return (
        <InlineField label="Legend"
                     tooltip="Controls the name of the time series, using name or pattern. For example {{resource}} will be replaced with label value for the label resource." labelWidth={14} interactive>
            <Input
                id="editor-query-legend-format"
                onChange={handleChange}
                value={value}
                width={80}
            />
        </InlineField>
    )
}
