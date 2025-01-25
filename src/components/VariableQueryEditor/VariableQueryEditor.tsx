import defaults from "lodash/defaults";
import React, {FC, useState} from "react";
import { InlineField, Select } from "@grafana/ui";
import { SelectableValue } from "@grafana/data";
import { Variable, VariableQuery, defaultVariableQuery } from "../../types";

export interface Props {
  query: VariableQuery;
  onChange: (query: VariableQuery, definition: string) => void;
}

export const VariableQueryEditor: FC<Props> = ({
  onChange,
  query: rawQuery,
}) => {
  const query = defaults(rawQuery, defaultVariableQuery);
  const [state, setState] = useState(query);

  const saveQuery = () => {
    onChange(state, `${state.selector.label}`);
  };

  const handleChange = (selector: SelectableValue<Variable>) =>
    setState({ ...state, selector });

  return (
      <InlineField label="Values for" labelWidth={16}>
          <Select
              width={16}
              maxVisibleValues={20}
              minMenuHeight={45}
              menuPlacement={"bottom"}
              onBlur={saveQuery}
              onChange={handleChange}
              value={state.selector}
              options={[
                  { value: Variable.Resource, label: "resourceID" },
                  { value: Variable.Vhost, label: "vhost" },
                  { value: Variable.Client, label: "client" },
                  { value: Variable.Region, label: "region" },
              ]}
          />
      </InlineField>
  );
};
