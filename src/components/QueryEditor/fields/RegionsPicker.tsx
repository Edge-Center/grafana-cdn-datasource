import React, {FC} from "react";
import {InlineField, Input} from "@grafana/ui";
import {useChangeInput} from "../useChangeInput";

export interface Props {
    value: string;
    onChange: (value: string) => void
}
export const RegionsPicker: FC<Props> = ({value, onChange}) => {
    const handleChange = useChangeInput(onChange);

    return (
        <InlineField label="Regions" tooltip="Filter by regions. Use commas to separate multiple values or reference variables." labelWidth={14} interactive>
            <Input
                id="editor-query-regions"
                onChange={handleChange}
                value={value}
                placeholder="Enter regions separated by commas"
                width={40}
            />
        </InlineField>
    )
}
