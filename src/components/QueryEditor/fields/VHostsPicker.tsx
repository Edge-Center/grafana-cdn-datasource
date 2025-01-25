import React, {FC} from "react";
import {InlineField, Input} from "@grafana/ui";
import {useChangeInput} from "../useChangeInput";

export interface Props {
    value: string;
    onChange: (value: string) => void
}
export const VHostsPicker: FC<Props> = ({value, onChange}) => {
    const handleChange = useChangeInput(onChange);

    return (
        <InlineField label="Vhosts" tooltip="Filter by vhosts. Use commas to separate multiple values or reference variables."
                     labelWidth={14} interactive>
            <Input
                id="editor-query-vhosts"
                onChange={handleChange}
                value={value}
                width={40}
                placeholder="Enter vhosts separated by commas"
            />
        </InlineField>
    )
}
